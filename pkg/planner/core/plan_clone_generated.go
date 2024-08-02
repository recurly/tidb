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

// Code generated by plan_clone_generator; DO NOT EDIT IT DIRECTLY.

package core

import (
	"github.com/pingcap/tidb/pkg/expression"
	"github.com/pingcap/tidb/pkg/planner/core/base"
	"github.com/pingcap/tidb/pkg/planner/util"
)

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalTableScan) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalTableScan)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	cloned.AccessCondition = util.CloneExpressions(op.AccessCondition)
	cloned.filterCondition = util.CloneExpressions(op.filterCondition)
	cloned.LateMaterializationFilterCondition = util.CloneExpressions(op.LateMaterializationFilterCondition)
	cloned.Ranges = util.CloneRanges(op.Ranges)
	cloned.HandleIdx = make([]int, len(op.HandleIdx))
	copy(cloned.HandleIdx, op.HandleIdx)
	if op.HandleCols != nil {
		cloned.HandleCols = op.HandleCols.Clone(newCtx.GetSessionVars().StmtCtx)
	}
	cloned.ByItems = util.CloneByItemss(op.ByItems)
	cloned.PlanPartInfo = op.PlanPartInfo.Clone()
	if op.SampleInfo != nil {
		return nil, false
	}
	cloned.constColsByCond = make([]bool, len(op.constColsByCond))
	copy(cloned.constColsByCond, op.constColsByCond)
	if op.runtimeFilterList != nil {
		return nil, false
	}
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexScan) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexScan)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	cloned.AccessCondition = util.CloneExpressions(op.AccessCondition)
	cloned.IdxCols = util.CloneColumns(op.IdxCols)
	cloned.IdxColLens = make([]int, len(op.IdxColLens))
	copy(cloned.IdxColLens, op.IdxColLens)
	cloned.Ranges = util.CloneRanges(op.Ranges)
	if op.GenExprs != nil {
		return nil, false
	}
	cloned.ByItems = util.CloneByItemss(op.ByItems)
	if op.pkIsHandleCol != nil {
		cloned.pkIsHandleCol = op.pkIsHandleCol.Clone().(*expression.Column)
	}
	cloned.constColsByCond = make([]bool, len(op.constColsByCond))
	copy(cloned.constColsByCond, op.constColsByCond)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalSelection) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalSelection)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalPlan.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalPlan = *basePlan
	cloned.Conditions = util.CloneExpressions(op.Conditions)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalProjection) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalProjection)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	cloned.Exprs = util.CloneExpressions(op.Exprs)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalSort) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalSort)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalPlan.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalPlan = *basePlan
	cloned.ByItems = util.CloneByItemss(op.ByItems)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalTopN) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalTopN)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalPlan.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalPlan = *basePlan
	cloned.ByItems = util.CloneByItemss(op.ByItems)
	cloned.PartitionBy = util.CloneSortItems(op.PartitionBy)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalStreamAgg) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalStreamAgg)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalAgg.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalAgg = *basePlan
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalHashAgg) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalHashAgg)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalAgg.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalAgg = *basePlan
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalHashJoin) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalHashJoin)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalJoin.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalJoin = *basePlan
	cloned.EqualConditions = util.CloneScalarFunctions(op.EqualConditions)
	cloned.NAEqualConditions = util.CloneScalarFunctions(op.NAEqualConditions)
	if op.runtimeFilterList != nil {
		return nil, false
	}
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalMergeJoin) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalMergeJoin)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalJoin.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalJoin = *basePlan
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalTableReader) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalTableReader)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	tablePlan, ok := op.tablePlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.tablePlan = tablePlan.(base.PhysicalPlan)
	cloned.TablePlans = flattenPushDownPlan(cloned.tablePlan)
	cloned.PlanPartInfo = op.PlanPartInfo.Clone()
	if op.TableScanAndPartitionInfos != nil {
		return nil, false
	}
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexReader) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexReader)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	indexPlan, ok := op.indexPlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.indexPlan = indexPlan.(base.PhysicalPlan)
	cloned.IndexPlans = flattenPushDownPlan(cloned.indexPlan)
	cloned.OutputColumns = util.CloneColumns(op.OutputColumns)
	cloned.PlanPartInfo = op.PlanPartInfo.Clone()
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PointGetPlan) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PointGetPlan)
	*cloned = *op
	cloned.Plan = *op.Plan.CloneWithNewCtx(newCtx)
	probeParents, ok := clonePhysicalPlansForPlanCache(newCtx, op.probeParents)
	if !ok {
		return nil, false
	}
	cloned.probeParents = probeParents
	cloned.PartitionNames = util.CloneCIStrs(op.PartitionNames)
	cloned.schema = op.schema.Clone()
	if op.PartitionIdx != nil {
		cloned.PartitionIdx = new(int)
		*cloned.PartitionIdx = *op.PartitionIdx
	}
	if op.Handle != nil {
		cloned.Handle = op.Handle.Copy()
	}
	if op.HandleConstant != nil {
		cloned.HandleConstant = op.HandleConstant.Clone().(*expression.Constant)
	}
	cloned.IndexValues = util.CloneDatums(op.IndexValues)
	cloned.IndexConstants = util.CloneConstants(op.IndexConstants)
	cloned.IdxCols = util.CloneColumns(op.IdxCols)
	cloned.IdxColLens = make([]int, len(op.IdxColLens))
	copy(cloned.IdxColLens, op.IdxColLens)
	cloned.AccessConditions = util.CloneExpressions(op.AccessConditions)
	cloned.ctx = newCtx
	cloned.accessCols = util.CloneColumns(op.accessCols)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *BatchPointGetPlan) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(BatchPointGetPlan)
	*cloned = *op
	cloned.baseSchemaProducer = *op.baseSchemaProducer.CloneWithNewCtx(newCtx)
	probeParents, ok := clonePhysicalPlansForPlanCache(newCtx, op.probeParents)
	if !ok {
		return nil, false
	}
	cloned.probeParents = probeParents
	cloned.PartitionNames = util.CloneCIStrs(op.PartitionNames)
	cloned.ctx = newCtx
	cloned.Handles = util.CloneHandles(op.Handles)
	cloned.HandleParams = util.CloneConstants(op.HandleParams)
	cloned.IndexValues = util.CloneDatum2D(op.IndexValues)
	cloned.IndexValueParams = util.CloneConstant2D(op.IndexValueParams)
	cloned.AccessConditions = util.CloneExpressions(op.AccessConditions)
	cloned.IdxCols = util.CloneColumns(op.IdxCols)
	cloned.IdxColLens = make([]int, len(op.IdxColLens))
	copy(cloned.IdxColLens, op.IdxColLens)
	cloned.PartitionIdxs = make([]int, len(op.PartitionIdxs))
	copy(cloned.PartitionIdxs, op.PartitionIdxs)
	cloned.accessCols = util.CloneColumns(op.accessCols)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalLimit) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalLimit)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	cloned.PartitionBy = util.CloneSortItems(op.PartitionBy)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexJoin) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexJoin)
	*cloned = *op
	basePlan, baseOK := op.basePhysicalJoin.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.basePhysicalJoin = *basePlan
	innerPlan, ok := op.innerPlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.innerPlan = innerPlan.(base.PhysicalPlan)
	cloned.Ranges = op.Ranges.CloneForPlanCache()
	cloned.KeyOff2IdxOff = make([]int, len(op.KeyOff2IdxOff))
	copy(cloned.KeyOff2IdxOff, op.KeyOff2IdxOff)
	cloned.IdxColLens = make([]int, len(op.IdxColLens))
	copy(cloned.IdxColLens, op.IdxColLens)
	if op.CompareFilters != nil {
		cloned.CompareFilters = op.CompareFilters.Copy()
	}
	cloned.OuterHashKeys = util.CloneColumns(op.OuterHashKeys)
	cloned.InnerHashKeys = util.CloneColumns(op.InnerHashKeys)
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexHashJoin) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexHashJoin)
	*cloned = *op
	inlj, ok := op.PhysicalIndexJoin.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.PhysicalIndexJoin = *inlj.(*PhysicalIndexJoin)
	cloned.self = cloned
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexLookUpReader) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexLookUpReader)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	indexPlan, ok := op.indexPlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.indexPlan = indexPlan.(base.PhysicalPlan)
	tablePlan, ok := op.tablePlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.tablePlan = tablePlan.(base.PhysicalPlan)
	cloned.IndexPlans = flattenPushDownPlan(cloned.indexPlan)
	cloned.TablePlans = flattenPushDownPlan(cloned.tablePlan)
	if op.ExtraHandleCol != nil {
		cloned.ExtraHandleCol = op.ExtraHandleCol.Clone().(*expression.Column)
	}
	cloned.PushedLimit = op.PushedLimit.Clone()
	cloned.CommonHandleCols = util.CloneColumns(op.CommonHandleCols)
	cloned.PlanPartInfo = op.PlanPartInfo.Clone()
	return cloned, true
}

// CloneForPlanCache implements the base.Plan interface.
func (op *PhysicalIndexMergeReader) CloneForPlanCache(newCtx base.PlanContext) (base.Plan, bool) {
	cloned := new(PhysicalIndexMergeReader)
	*cloned = *op
	basePlan, baseOK := op.physicalSchemaProducer.cloneForPlanCacheWithSelf(newCtx, cloned)
	if !baseOK {
		return nil, false
	}
	cloned.physicalSchemaProducer = *basePlan
	cloned.PushedLimit = op.PushedLimit.Clone()
	cloned.ByItems = util.CloneByItemss(op.ByItems)
	partialPlans, ok := clonePhysicalPlansForPlanCache(newCtx, op.partialPlans)
	if !ok {
		return nil, false
	}
	cloned.partialPlans = partialPlans
	tablePlan, ok := op.tablePlan.CloneForPlanCache(newCtx)
	if !ok {
		return nil, false
	}
	cloned.tablePlan = tablePlan.(base.PhysicalPlan)
	cloned.PartialPlans = make([][]base.PhysicalPlan, len(op.PartialPlans))
	for i, plan := range cloned.partialPlans {
		cloned.PartialPlans[i] = flattenPushDownPlan(plan)
	}
	cloned.TablePlans = flattenPushDownPlan(cloned.tablePlan)
	cloned.PlanPartInfo = op.PlanPartInfo.Clone()
	if op.HandleCols != nil {
		cloned.HandleCols = op.HandleCols.Clone(newCtx.GetSessionVars().StmtCtx)
	}
	return cloned, true
}
