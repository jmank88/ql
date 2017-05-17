package ql

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCQL(t *testing.T) {
	const query = `SELECT a, b, c, d FROM target WHERE (a = ? AND (b < ? OR c != ?) AND d > ?) ORDER BY a DESC LIMIT 10`
	q, a := CQL(&complexQuery)
	if q != query {
		t.Errorf("expected %q but got %q", query, q)
	}
	var args = []interface{}{"test", json.Number("2"), json.Number("5"), json.Number("0")}
	if !reflect.DeepEqual(args, a) {
		t.Errorf("expected %v but got %v", args, a)
	}
}
