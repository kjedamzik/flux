package transformations

import (
	"fmt"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/plan"
)

const ColumnsKind = "columns"

type ColumnsOpSpec struct {
}

func init() {
	keysSignature := flux.FunctionSignature(nil, nil)

	flux.RegisterFunction(ColumnsKind, createColumnsOpSpec, keysSignature)
	flux.RegisterOpSpec(ColumnsKind, newColumnsOp)
	plan.RegisterProcedureSpec(ColumnsKind, newColumnsProcedure, ColumnsKind)
	execute.RegisterTransformation(ColumnsKind, createColumnsTransformation)
}

func createColumnsOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	if err := a.AddParentFromArgs(args); err != nil {
		return nil, err
	}

	return new(ColumnsOpSpec), nil
}

func newColumnsOp() flux.OperationSpec {
	return new(ColumnsOpSpec)
}

func (s *ColumnsOpSpec) Kind() flux.OperationKind {
	return ColumnsKind
}

type ColumnsProcedureSpec struct {
	plan.DefaultCost
}

func newColumnsProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	_, ok := qs.(*ColumnsOpSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}

	return &ColumnsProcedureSpec{}, nil
}

func (s *ColumnsProcedureSpec) Kind() plan.ProcedureKind {
	return ColumnsKind
}

func (s *ColumnsProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(ColumnsProcedureSpec)

	*ns = *s

	return ns
}

func createColumnsTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration) (execute.Transformation, execute.Dataset, error) {
	s, ok := spec.(*ColumnsProcedureSpec)
	if !ok {
		return nil, nil, fmt.Errorf("invalid spec type %T", spec)
	}
	cache := execute.NewTableBuilderCache(a.Allocator())
	d := execute.NewDataset(id, mode, cache)
	t := NewColumnsTransformation(d, cache, s)
	return t, d, nil
}

type columnsTransformation struct {
	d     execute.Dataset
	cache execute.TableBuilderCache
}

func NewColumnsTransformation(d execute.Dataset, cache execute.TableBuilderCache, spec *ColumnsProcedureSpec) *columnsTransformation {
	return &columnsTransformation{
		d:     d,
		cache: cache,
	}
}

func (t *columnsTransformation) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return t.d.RetractTable(key)
}

func (t *columnsTransformation) Process(id execute.DatasetID, tbl flux.Table) error {
	builder, created := t.cache.TableBuilder(tbl.Key())
	if !created {
		return fmt.Errorf("columns found duplicate table with key: %v", tbl.Key())
	}

	labels := make([]string, len(tbl.Cols()))
	for i, c := range tbl.Cols() {
		labels[i] = c.Label
	}

	// Add the group key columns to this table.
	if err := execute.AddTableKeyCols(tbl.Key(), builder); err != nil {
		return err
	}

	// Create a new column for the column names.
	colIdx, err := builder.AddCol(flux.ColMeta{Label: execute.DefaultValueColLabel, Type: flux.TString})
	if err != nil {
		return err
	}

	// Append the key values repeatedly to the table.
	for i := 0; i < len(labels); i++ {
		if err := execute.AppendKeyValues(tbl.Key(), builder); err != nil {
			return err
		}
	}

	// Append the labels to the column index.
	if err := builder.AppendStrings(colIdx, labels); err != nil {
		return err
	}

	return nil
}

func (t *columnsTransformation) UpdateWatermark(id execute.DatasetID, mark execute.Time) error {
	return t.d.UpdateWatermark(mark)
}
func (t *columnsTransformation) UpdateProcessingTime(id execute.DatasetID, pt execute.Time) error {
	return t.d.UpdateProcessingTime(pt)
}
func (t *columnsTransformation) Finish(id execute.DatasetID, err error) {
	t.d.Finish(err)
}
