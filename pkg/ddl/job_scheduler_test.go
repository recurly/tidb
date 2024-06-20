// Copyright 2024 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ddl

import (
	"context"
	"testing"
	"time"

	"github.com/pingcap/errors"
	"github.com/pingcap/tidb/pkg/ddl/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func reduceIntervals(t testing.TB) {
	loopRetryIntBak := schedulerLoopRetryInterval
	schedulerLoopRetryInterval = 10 * time.Millisecond
	t.Cleanup(func() {
		schedulerLoopRetryInterval = loopRetryIntBak
	})
}

func TestMustReloadSchemas(t *testing.T) {
	reduceIntervals(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	loader := mock.NewMockSchemaLoader(ctrl)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sch := &jobScheduler{
		schCtx:       ctx,
		schemaLoader: loader,
	}
	// directly success
	loader.EXPECT().Reload().Return(nil)
	sch.mustReloadSchemas()
	require.True(t, ctrl.Satisfied())
	// success after retry
	loader.EXPECT().Reload().Return(errors.New("mock err"))
	loader.EXPECT().Reload().Return(nil)
	sch.mustReloadSchemas()
	require.True(t, ctrl.Satisfied())
	// exit on context cancel
	loader.EXPECT().Reload().Do(func() error {
		cancel()
		return errors.New("mock err")
	})
	sch.mustReloadSchemas()
	require.True(t, ctrl.Satisfied())
}

func TestRefreshMinJobID(t *testing.T) {
	reduceIntervals(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mgr := mock.NewMockManager(ctrl)

	sch := &jobScheduler{
		schCtx:    context.Background(),
		sysTblMgr: mgr,
	}
	// success
	start := time.Now()
	mgr.EXPECT().GetMinJobID(gomock.Any(), int64(0)).Return(int64(1), nil)
	sch.refreshMinJobID()
	require.EqualValues(t, 1, sch.currMinJobID)
	require.GreaterOrEqual(t, sch.lastRefreshMinIDTime, start)
	require.True(t, ctrl.Satisfied())
	// not refresh too fast
	schedulerLoopRetryInterval = time.Hour
	sch.refreshMinJobID()
	require.True(t, ctrl.Satisfied())
	// ignore refresh error
	sch.lastRefreshMinIDTime = time.Time{}
	mgr.EXPECT().GetMinJobID(gomock.Any(), int64(1)).Return(int64(0), errors.New("mock err"))
	sch.refreshMinJobID()
	require.EqualValues(t, 1, sch.currMinJobID)
	require.True(t, ctrl.Satisfied())
	// success again
	sch.lastRefreshMinIDTime = time.Time{}
	mgr.EXPECT().GetMinJobID(gomock.Any(), int64(1)).Return(int64(100), nil)
	sch.refreshMinJobID()
	require.EqualValues(t, 100, sch.currMinJobID)
	require.GreaterOrEqual(t, sch.lastRefreshMinIDTime, start)
	require.True(t, ctrl.Satisfied())
}
