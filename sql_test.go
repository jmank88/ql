package ql

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPOSTGRESQL(t *testing.T) {
	const query = `SELECT a, b, c, d FROM target WHERE (a = $1 AND (b < $2 OR c != $3) AND d > $4) ORDER BY a DESC LIMIT 10`
	q, a := POSTGRESQL(&complexQuery)
	if q != query {
		t.Errorf("expected %q but got %q", query, q)
	}
	var args = []interface{}{"test", json.Number("2"), json.Number("5"), json.Number("0")}
	if !reflect.DeepEqual(args, a) {
		t.Errorf("expected %v but got %v", args, a)
	}
}

func TestMYSQL(t *testing.T) {
	const query = `SELECT a, b, c, d FROM target WHERE (a = ? AND (b < ? OR c != ?) AND d > ?) ORDER BY a DESC LIMIT 10`
	q, a := MYSQL(&complexQuery)
	if q != query {
		t.Errorf("expected %q but got %q", query, q)
	}
	var args = []interface{}{"test", json.Number("2"), json.Number("5"), json.Number("0")}
	if !reflect.DeepEqual(args, a) {
		t.Errorf("expected %v but got %v", args, a)
	}
}
