package transformations

import (
	"fmt"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/semantic"
)

const KeysKind = "keys"

type KeysOpSpec struct {
}

func init() {
	keysSignature := flux.FunctionSignature(
		map[string]semantic.PolyType{
			"except": semantic.NewArrayPolyType(semantic.String),
		},
		nil,
	)

	flux.RegisterFunction(KeysKind, createKeysOpSpec, keysSignature)
	flux.RegisterOpSpec(KeysKind, newKeysOp)
	plan.RegisterProcedureSpec(KeysKind, newKeysProcedure, KeysKind)
	execute.RegisterTransformation(KeysKind, createKeysTransformation)
}

func createKeysOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	if err := a.AddParentFromArgs(args); err != nil {
		return nil, err
	}

	return new(KeysOpSpec), nil
}

func newKeysOp() flux.OperationSpec {
	return new(KeysOpSpec)
}

func (s *KeysOpSpec) Kind() flux.OperationKind {
	return KeysKind
}

type KeysProcedureSpec struct {
	plan.DefaultCost
}

func newKeysProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	_, ok := qs.(*KeysOpSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}

	return &KeysProcedureSpec{}, nil
}

func (s *KeysProcedureSpec) Kind() plan.ProcedureKind {
	return KeysKind
}

func (s *KeysProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(KeysProcedureSpec)

	*ns = *s

	return ns
}

func createKeysTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration) (execute.Transformation, execute.Dataset, error) {
	s, ok := spec.(*KeysProcedureSpec)
	if !ok {
		return nil, nil, fmt.Errorf("invalid spec type %T", spec)
	}
	cache := execute.NewTableBuilderCache(a.Allocator())
	d := execute.NewDataset(id, mode, cache)
	t := NewKeysTransformation(d, cache, s)
	return t, d, nil
}

type keysTransformation struct {
	d     execute.Dataset
	cache execute.TableBuilderCache
}

func NewKeysTransformation(d execute.Dataset, cache execute.TableBuilderCache, spec *KeysProcedureSpec) *keysTransformation {
	return &keysTransformation{
		d:     d,
		cache: cache,
	}
}

func (t *keysTransformation) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return t.d.RetractTable(key)
}

func (t *keysTransformation) Process(id execute.DatasetID, tbl flux.Table) error {
	builder, created := t.cache.TableBuilder(tbl.Key())
	if !created {
		return fmt.Errorf("keys found duplicate table with key: %v", tbl.Key())
	}

	keys := make([]string, 0, len(tbl.Cols()))
	for _, c := range tbl.Key().Cols() {
		keys = append(keys, c.Label)
	}

	// Add the key to this table.
	if err := execute.AddTableKeyCols(tbl.Key(), builder); err != nil {
		return err
	}

	colIdx, err := builder.AddCol(flux.ColMeta{Label: execute.DefaultValueColLabel, Type: flux.TString})
	if err != nil {
		return err
	}

	// Append the key values repeatedly to the table.
	for i := 0; i < len(keys); i++ {
		if err := execute.AppendKeyValues(tbl.Key(), builder); err != nil {
			return err
		}
	}

	// Append the keys to the column index.
	if err := builder.AppendStrings(colIdx, keys); err != nil {
		return err
	}

	return nil
}

func (t *keysTransformation) UpdateWatermark(id execute.DatasetID, mark execute.Time) error {
	return t.d.UpdateWatermark(mark)
}
func (t *keysTransformation) UpdateProcessingTime(id execute.DatasetID, pt execute.Time) error {
	return t.d.UpdateProcessingTime(pt)
}
func (t *keysTransformation) Finish(id execute.DatasetID, err error) {
	t.d.Finish(err)
}
