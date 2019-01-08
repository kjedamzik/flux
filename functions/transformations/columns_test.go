package transformations_test

import (
	"testing"
	"time"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/execute/executetest"
	"github.com/influxdata/flux/functions/inputs"
	"github.com/influxdata/flux/functions/transformations"
	"github.com/influxdata/flux/querytest"
)

func TestColumns_NewQuery(t *testing.T) {
	tests := []querytest.NewQueryTestCase{
		{
			Name: "from range columns",
			Raw:  `from(bucket:"mydb") |> range(start:-1h) |> columns()`,
			Want: &flux.Spec{
				Operations: []*flux.Operation{
					{
						ID: "from0",
						Spec: &inputs.FromOpSpec{
							Bucket: "mydb",
						},
					},
					{
						ID: "range1",
						Spec: &transformations.RangeOpSpec{
							Start: flux.Time{
								Relative:   -1 * time.Hour,
								IsRelative: true,
							},
							Stop:        flux.Now,
							TimeColumn:  "_time",
							StartColumn: "_start",
							StopColumn:  "_stop",
						},
					},
					{
						ID: "columns2",
						Spec: &transformations.ColumnsOpSpec{
						},
					},
				},
				Edges: []flux.Edge{
					{Parent: "from0", Child: "range1"},
					{Parent: "range1", Child: "columns2"},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			querytest.NewQueryTestHelper(t, tc)
		})
	}
}

func TestColumnsOperation_Marshaling(t *testing.T) {
	data := []byte(`{"id":"columns","kind":"columns","spec":{}}`)
	op := &flux.Operation{
		ID:   "columns",
		Spec: &transformations.ColumnsOpSpec{},
	}

	querytest.OperationMarshalingTestHelper(t, data, op)
}

func TestColumns_Process(t *testing.T) {
	spec := &transformations.ColumnsProcedureSpec{}

	testCases := []struct {
		name string
		data []flux.Table
		want []*executetest.Table
	}{
		{
			name: "one table",
			data: []flux.Table{
				&executetest.Table{
					KeyCols: []string{"tag0", "tag1"},
					ColMeta: []flux.ColMeta{
						{Label: "_time", Type: flux.TTime},
						{Label: "_value", Type: flux.TFloat},
						{Label: "tag0", Type: flux.TString},
						{Label: "tag1", Type: flux.TString},
						{Label: "valid", Type: flux.TBool},
					},
					Data: [][]interface{}{
						{execute.Time(1), 2.0, "a", "b", true},
					},
				},
			},
			want: []*executetest.Table{{
				KeyCols: []string{"tag0", "tag1"},
				ColMeta: []flux.ColMeta{
					{Label: "tag0", Type: flux.TString},
					{Label: "tag1", Type: flux.TString},
					{Label: "_value", Type: flux.TString},
				},
				Data: [][]interface{}{
					{"a", "b", "_time"},
					{"a", "b", "_value"},
					{"a", "b", "tag0"},
					{"a", "b", "tag1"},
					{"a", "b", "valid"},
				},
			}},
		},
		{
			name: "three tables",
			data: []flux.Table{
				&executetest.Table{
					KeyCols: []string{"tag0", "tag1"},
					ColMeta: []flux.ColMeta{
						{Label: "_time", Type: flux.TTime},
						{Label: "_value", Type: flux.TFloat},
						{Label: "tag0", Type: flux.TString},
						{Label: "tag1", Type: flux.TString},
					},
					Data: [][]interface{}{
						{execute.Time(1), 2.0, "a", "b"},
					},
				},
				&executetest.Table{
					KeyCols: []string{"tag0", "tag1"},
					ColMeta: []flux.ColMeta{
						{Label: "_time", Type: flux.TTime},
						{Label: "_value", Type: flux.TFloat},
						{Label: "tag0", Type: flux.TString},
						{Label: "tag1", Type: flux.TString},
					},
					Data: [][]interface{}{
						{execute.Time(1), 2.0, "b", "b"},
					},
				},
				&executetest.Table{
					KeyCols: []string{"tag0", "tag2"},
					ColMeta: []flux.ColMeta{
						{Label: "_time", Type: flux.TTime},
						{Label: "_value", Type: flux.TFloat},
						{Label: "tag0", Type: flux.TString},
						{Label: "tag2", Type: flux.TString},
					},
					Data: [][]interface{}{
						{execute.Time(1), 2.0, "b", "c",},
					},
				},
			},
			want: []*executetest.Table{
				{
					KeyCols: []string{"tag0", "tag1"},
					ColMeta: []flux.ColMeta{
						{Label: "tag0", Type: flux.TString},
						{Label: "tag1", Type: flux.TString},
						{Label: "_value", Type: flux.TString},
					},
					Data: [][]interface{}{
						{"a", "b", "_time"},
						{"a", "b", "_value"},
						{"a", "b", "tag0"},
						{"a", "b", "tag1"},
					},
				},
				{
					KeyCols: []string{"tag0", "tag1"},
					ColMeta: []flux.ColMeta{
						{Label: "tag0", Type: flux.TString},
						{Label: "tag1", Type: flux.TString},
						{Label: "_value", Type: flux.TString},
					},
					Data: [][]interface{}{
						{"b", "b", "_time"},
						{"b", "b", "_value"},
						{"b", "b", "tag0"},
						{"b", "b", "tag1"},
					},
				},
				{
					KeyCols: []string{"tag0", "tag2"},
					ColMeta: []flux.ColMeta{
						{Label: "tag0", Type: flux.TString},
						{Label: "tag2", Type: flux.TString},
						{Label: "_value", Type: flux.TString},
					},
					Data: [][]interface{}{
						{"b", "c", "_time"},
						{"b", "c", "_value"},
						{"b", "c", "tag0"},
						{"b", "c", "tag2"},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			executetest.ProcessTestHelper(
				t,
				tc.data,
				tc.want,
				nil,
				func(d execute.Dataset, c execute.TableBuilderCache) execute.Transformation {
					return transformations.NewColumnsTransformation(d, c, spec)
				},
			)
		})
	}
}
