package executetest

import (
	"fmt"

	"github.com/apache/arrow/go/arrow/array"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/arrow"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
)

// Table is an implementation of execute.Table
// It is designed to make it easy to statically declare the data within the table.
// Not all fields need to be set. See comments on each field.
// Use Normalize to ensure that all fields are set before equality comparisons.
type Table struct {
	// GroupKey of the table. Does not need to be set explicitly.
	GroupKey flux.GroupKey
	// KeyCols is a list of column that are part of the group key.
	// The column type is deduced from the ColMeta slice.
	KeyCols []string
	// KeyValues is a list of values for the group key columns.
	// Only needs to be set when no data is present on the table.
	KeyValues []interface{}
	// ColMeta is a list of columns of the table.
	ColMeta []flux.ColMeta
	// Data is a list of rows, i.e. Data[row][col]
	// Each row must be a list with length equal to len(ColMeta)
	Data [][]interface{}
}

// Normalize ensures all fields of the table are set correctly.
func (t *Table) Normalize() {
	if t.GroupKey == nil {
		cols := make([]flux.ColMeta, len(t.KeyCols))
		vs := make([]values.Value, len(t.KeyCols))
		if len(t.KeyValues) != len(t.KeyCols) {
			t.KeyValues = make([]interface{}, len(t.KeyCols))
		}
		for j, label := range t.KeyCols {
			idx := execute.ColIdx(label, t.ColMeta)
			if idx < 0 {
				panic(fmt.Errorf("table invalid: missing group column %q", label))
			}
			cols[j] = t.ColMeta[idx]
			if len(t.Data) > 0 {
				t.KeyValues[j] = t.Data[0][idx]
			}
			v := values.New(t.KeyValues[j])
			if v.Type() == semantic.Invalid {
				panic(fmt.Errorf("invalid value: %s", t.KeyValues[j]))
			}
			vs[j] = v
		}
		t.GroupKey = execute.NewGroupKey(cols, vs)
	}
}

func (t *Table) Empty() bool {
	return len(t.Data) == 0
}

func (t *Table) RefCount(n int) {}

func (t *Table) Cols() []flux.ColMeta {
	return t.ColMeta
}

func (t *Table) Key() flux.GroupKey {
	t.Normalize()
	return t.GroupKey
}

func (t *Table) Do(f func(flux.ColReader) error) error {
	for _, r := range t.Data {
		if err := f(ColReader{
			key:  t.Key(),
			cols: t.ColMeta,
			row:  r,
		}); err != nil {
			return err
		}
	}
	return nil
}

// RowWiseArrowTable is a flux Table implementation that
// calls f once for each row in its DoArrow method.
type RowWiseArrowTable struct {
	*Table
}

// DoArrow calls f once for each row in the table
func (t *RowWiseArrowTable) DoArrow(f func(flux.ArrowColReader) error) error {
	cols := make([]array.Interface, len(t.ColMeta))
	for j, col := range t.ColMeta {
		switch col.Type {
		case flux.TBool:
			b := arrow.NewBoolBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(bool))
			}
			cols[j] = b.NewBooleanArray()
			b.Release()
		case flux.TFloat:
			b := arrow.NewFloatBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(float64))
			}
			cols[j] = b.NewFloat64Array()
			b.Release()
		case flux.TInt:
			b := arrow.NewIntBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(int64))
			}
			cols[j] = b.NewInt64Array()
			b.Release()
		case flux.TString:
			b := arrow.NewStringBuilder(nil)
			for i := range t.Data {
				b.AppendString(t.Data[i][j].(string))
			}
			cols[j] = b.NewBinaryArray()
			b.Release()
		case flux.TTime:
			b := arrow.NewIntBuilder(nil)
			for i := range t.Data {
				b.Append(int64(t.Data[i][j].(values.Time)))
			}
			cols[j] = b.NewInt64Array()
			b.Release()
		case flux.TUInt:
			b := arrow.NewUintBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(uint64))
			}
			cols[j] = b.NewUint64Array()
			b.Release()
		}
	}

	release := func(cols []array.Interface) {
		for _, arr := range cols {
			arr.Release()
		}
	}
	defer release(cols)

	l := cols[0].Len()
	for i := 0; i < l; i++ {
		row := make([]array.Interface, len(t.ColMeta))
		for j, col := range t.ColMeta {
			switch col.Type {
			case flux.TBool:
				row[j] = arrow.BoolSlice(cols[j].(*array.Boolean), i, i+1)
			case flux.TFloat:
				row[j] = arrow.FloatSlice(cols[j].(*array.Float64), i, i+1)
			case flux.TInt:
				row[j] = arrow.IntSlice(cols[j].(*array.Int64), i, i+1)
			case flux.TString:
				row[j] = arrow.StringSlice(cols[j].(*array.Binary), i, i+1)
			case flux.TTime:
				row[j] = arrow.IntSlice(cols[j].(*array.Int64), i, i+1)
			case flux.TUInt:
				row[j] = arrow.UintSlice(cols[j].(*array.Uint64), i, i+1)
			}
		}
		if err := f(&ArrowColReader{
			key:  t.Key(),
			meta: t.ColMeta,
			cols: row,
		}); err != nil {
			return err
		}
		release(row)
	}
	return nil
}

func (t *Table) DoArrow(f func(flux.ArrowColReader) error) error {
	cols := make([]array.Interface, len(t.ColMeta))
	for j, col := range t.ColMeta {
		switch col.Type {
		case flux.TBool:
			b := arrow.NewBoolBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(bool))
			}
			cols[j] = b.NewBooleanArray()
			b.Release()
		case flux.TFloat:
			b := arrow.NewFloatBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(float64))
			}
			cols[j] = b.NewFloat64Array()
			b.Release()
		case flux.TInt:
			b := arrow.NewIntBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(int64))
			}
			cols[j] = b.NewInt64Array()
			b.Release()
		case flux.TString:
			b := arrow.NewStringBuilder(nil)
			for i := range t.Data {
				b.AppendString(t.Data[i][j].(string))
			}
			cols[j] = b.NewBinaryArray()
			b.Release()
		case flux.TTime:
			b := arrow.NewIntBuilder(nil)
			for i := range t.Data {
				b.Append(int64(t.Data[i][j].(values.Time)))
			}
			cols[j] = b.NewInt64Array()
			b.Release()
		case flux.TUInt:
			b := arrow.NewUintBuilder(nil)
			for i := range t.Data {
				b.Append(t.Data[i][j].(uint64))
			}
			cols[j] = b.NewUint64Array()
			b.Release()
		}
	}

	cr := &ArrowColReader{
		key:  t.Key(),
		meta: t.ColMeta,
		cols: cols,
	}
	return f(cr)
}

func (t *Table) Statistics() flux.Statistics { return flux.Statistics{} }

type ColReader struct {
	key  flux.GroupKey
	cols []flux.ColMeta
	row  []interface{}
}

func (cr ColReader) Cols() []flux.ColMeta {
	return cr.cols
}

func (cr ColReader) Key() flux.GroupKey {
	return cr.key
}
func (cr ColReader) Len() int {
	return 1
}

func (cr ColReader) Bools(j int) []bool {
	return []bool{cr.row[j].(bool)}
}

func (cr ColReader) Ints(j int) []int64 {
	return []int64{cr.row[j].(int64)}
}

func (cr ColReader) UInts(j int) []uint64 {
	return []uint64{cr.row[j].(uint64)}
}

func (cr ColReader) Floats(j int) []float64 {
	return []float64{cr.row[j].(float64)}
}

func (cr ColReader) Strings(j int) []string {
	return []string{cr.row[j].(string)}
}

func (cr ColReader) Times(j int) []execute.Time {
	return []execute.Time{cr.row[j].(execute.Time)}
}

type ArrowColReader struct {
	key  flux.GroupKey
	meta []flux.ColMeta
	cols []array.Interface
}

func (cr *ArrowColReader) Key() flux.GroupKey {
	return cr.key
}

func (cr *ArrowColReader) Cols() []flux.ColMeta {
	return cr.meta
}

func (cr *ArrowColReader) Len() int {
	if len(cr.cols) == 0 {
		return 0
	}
	return cr.cols[0].Len()
}

func (cr *ArrowColReader) Bools(j int) *array.Boolean {
	return cr.cols[j].(*array.Boolean)
}

func (cr *ArrowColReader) Ints(j int) *array.Int64 {
	return cr.cols[j].(*array.Int64)
}

func (cr *ArrowColReader) UInts(j int) *array.Uint64 {
	return cr.cols[j].(*array.Uint64)
}

func (cr *ArrowColReader) Floats(j int) *array.Float64 {
	return cr.cols[j].(*array.Float64)
}

func (cr *ArrowColReader) Strings(j int) *array.Binary {
	return cr.cols[j].(*array.Binary)
}

func (cr *ArrowColReader) Times(j int) *array.Int64 {
	return cr.cols[j].(*array.Int64)
}

func TablesFromCache(c execute.DataCache) (tables []*Table, err error) {
	c.ForEach(func(key flux.GroupKey) {
		if err != nil {
			return
		}
		var tbl flux.Table
		tbl, err = c.Table(key)
		if err != nil {
			return
		}
		var cb *Table
		cb, err = ConvertTable(tbl)
		if err != nil {
			return
		}
		tables = append(tables, cb)
		c.ExpireTable(key)
	})
	return tables, nil
}

func ConvertTable(tbl flux.Table) (*Table, error) {
	key := tbl.Key()
	blk := &Table{
		GroupKey: key,
		ColMeta:  tbl.Cols(),
	}

	keyCols := key.Cols()
	if len(keyCols) > 0 {
		blk.KeyCols = make([]string, len(keyCols))
		blk.KeyValues = make([]interface{}, len(keyCols))
		for j, c := range keyCols {
			blk.KeyCols[j] = c.Label
			var v interface{}
			switch c.Type {
			case flux.TBool:
				v = key.ValueBool(j)
			case flux.TUInt:
				v = key.ValueUInt(j)
			case flux.TInt:
				v = key.ValueInt(j)
			case flux.TFloat:
				v = key.ValueFloat(j)
			case flux.TString:
				v = key.ValueString(j)
			case flux.TTime:
				v = key.ValueTime(j)
			default:
				return nil, fmt.Errorf("unsupported column type %v", c.Type)
			}
			blk.KeyValues[j] = v
		}
	}

	err := tbl.Do(func(cr flux.ColReader) error {
		l := cr.Len()
		for i := 0; i < l; i++ {
			row := make([]interface{}, len(blk.ColMeta))
			for j, c := range blk.ColMeta {
				var v interface{}
				switch c.Type {
				case flux.TBool:
					v = cr.Bools(j)[i]
				case flux.TInt:
					v = cr.Ints(j)[i]
				case flux.TUInt:
					v = cr.UInts(j)[i]
				case flux.TFloat:
					v = cr.Floats(j)[i]
				case flux.TString:
					v = cr.Strings(j)[i]
				case flux.TTime:
					v = cr.Times(j)[i]
				default:
					panic(fmt.Errorf("unknown column type %s", c.Type))
				}
				row[j] = v
			}
			blk.Data = append(blk.Data, row)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return blk, nil
}

type SortedTables []*Table

func (b SortedTables) Len() int {
	return len(b)
}

func (b SortedTables) Less(i int, j int) bool {
	return b[i].Key().Less(b[j].Key())
}

func (b SortedTables) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

// NormalizeTables ensures that each table is normalized
func NormalizeTables(bs []*Table) {
	for _, b := range bs {
		b.Key()
	}
}

func MustCopyTable(tbl flux.Table) flux.Table {
	cpy, _ := execute.CopyTable(tbl, UnlimitedAllocator)
	return cpy
}
