package ql

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	b, err := json.Marshal(&complexQuery)
	if err != nil {
		t.Fatal(err)
	}
	var v Query
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err = d.Decode(&v)
	if err != nil {
		t.Fatal(err)
	}

	v.checkEqual(t, &complexQuery)
}

func (q *Query) checkEqual(t *testing.T, q2 *Query) {
	if len(q.Fields) != len(q2.Fields) {
		t.Error("different number of fields")
	} else {
		for i := range q.Fields {
			if q.Fields[i] != q2.Fields[i] {
				t.Error("different fields")
			}
		}
	}
	if q.Filter == nil {
		if q2.Filter != nil {
			t.Error("nil and non-nil filter")
		}
	} else {
		q.Filter.checkEqual(t, q2.Filter)
	}
	if len(q.Sort) != len(q2.Sort) {
		t.Error("different number of sorts")
	}
	for i := range q.Sort {
		if q.Sort[i] != q2.Sort[i] {
			t.Error("different sorts")
		}
	}
	if q.Max == nil {
		if q2.Max != nil {
			t.Error("nil and non-nil sort")
		}
	} else {
		if *q.Max != *q2.Max {
			t.Error("different maximums")
		}
	}
}

func (c *Filter) checkEqual(t *testing.T, c2 *Filter) {
	if c.Comparator != c2.Comparator {
		t.Error("different comparators")
	}
	if c.Field != c2.Field {
		t.Error("different fields")
	}
	if !reflect.DeepEqual(c.Value, c2.Value) {
		t.Errorf("different values: %T %v != %T %v", c.Value, c.Value, c2.Value, c2.Value)
	}
	if c.Operator != c2.Operator {
		t.Error("different operators")
	}
	if len(c.Filters) != len(c2.Filters) {
		t.Error("different number of filters")
	}
	for i := range c.Filters {
		c.Filters[i].checkEqual(t, c2.Filters[i])
	}
}

var complexQuery = Query{
	Fields: []string{"a", "b", "c", "d"},
	Target: "target",
	Filter: And(
		Compare("a", "=", "test"),
		Or(Compare("b", "<", json.Number("2")),
			Compare("c", "!=", json.Number("5"))),
		Compare("d", ">", json.Number("0")),
	),
	Sort: []Sort{
		Desc("a"),
	},
	Max: newInt(10),
}

func newInt(i int) *int {
	return &i
}
