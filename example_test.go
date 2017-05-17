package ql_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmank88/ql"
)

func ExampleQuery() {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	_ = e.Encode(&ql.Query{
		Fields: []string{"a", "b", "c", "d"},
		Target: "target",
		Filter: ql.And(
			ql.Compare("a", "=", "test"),
			ql.Or(
				ql.Compare("b", "<", json.Number("2")),
				ql.Compare("c", "!=", json.Number("5")),
			),
			ql.Compare("d", ">", json.Number("0")),
		),
		Sort: []ql.Sort{
			{
				Field:      "a",
				Descending: true,
			},
		},
	})

	// Output:
	// {
	// 	"target": "target",
	// 	"fields": [
	// 		"a",
	// 		"b",
	// 		"c",
	// 		"d"
	// 	],
	// 	"filter": {
	// 		"operator": "AND",
	// 		"filters": [
	// 			{
	// 				"comparator": "=",
	// 				"field": "a",
	// 				"value": "test"
	// 			},
	// 			{
	// 				"operator": "OR",
	// 				"filters": [
	// 					{
	// 						"comparator": "\u003c",
	// 						"field": "b",
	// 						"value": 2
	// 					},
	// 					{
	// 						"comparator": "!=",
	// 						"field": "c",
	// 						"value": 5
	// 					}
	// 				]
	// 			},
	// 			{
	// 				"comparator": "\u003e",
	// 				"field": "d",
	// 				"value": 0
	// 			}
	// 		]
	// 	},
	// 	"sort": [
	// 		{
	// 			"field": "a",
	// 			"descending": true
	// 		}
	// 	]
	// }
}

func ExamplePOSTGRESQL() {
	q := &ql.Query{
		Fields: []string{"a", "b", "c", "d"},
		Target: "target",
		Filter: ql.And(
			ql.Compare("a", "=", "test"),
			ql.Or(
				ql.Compare("b", "<", 2),
				ql.Compare("c", "!=", 5),
			),
			ql.Compare("d", ">", 0),
		),
		Sort: []ql.Sort{
			{
				Field:      "a",
				Descending: true,
			},
		},
		Max: newInt(100),
	}
	query, args := ql.POSTGRESQL(q)
	// db.Query(query, args...)
	fmt.Println(query)
	fmt.Println(args)

	// Output:
	// SELECT a, b, c, d FROM target WHERE (a = $1 AND (b < $2 OR c != $3) AND d > $4) ORDER BY a DESC LIMIT 100
	// [test 2 5 0]
}

func ExampleMYSQL() {
	q := &ql.Query{
		Fields: []string{"a", "b", "c", "d"},
		Target: "target",
		Filter: ql.And(
			ql.Compare("a", "=", "test"),
			ql.Or(
				ql.Compare("b", "<", 2),
				ql.Compare("c", "!=", 5),
			),
			ql.Compare("d", ">", 0),
		),
		Sort: []ql.Sort{
			{
				Field:      "a",
				Descending: true,
			},
		},
		Max: newInt(100),
	}
	query, args := ql.MYSQL(q)
	// db.Query(query, args...)
	fmt.Println(query)
	fmt.Println(args)

	// Output:
	// SELECT a, b, c, d FROM target WHERE (a = ? AND (b < ? OR c != ?) AND d > ?) ORDER BY a DESC LIMIT 100
	// [test 2 5 0]
}

func ExampleCQL() {
	q := &ql.Query{
		Fields: []string{"a", "b", "c", "d"},
		Target: "target",
		Filter: ql.And(
			ql.Compare("a", "=", "test"),
			ql.Or(
				ql.Compare("b", "<", 2),
				ql.Compare("c", "!=", 5),
			),
			ql.Compare("d", ">", 0),
		),
		Sort: []ql.Sort{
			{
				Field:      "a",
				Descending: true,
			},
		},
		Max: newInt(20),
	}
	query, args := ql.MYSQL(q)
	// session.Query(query, args...)
	fmt.Println(query)
	fmt.Println(args)

	// Output:
	// SELECT a, b, c, d FROM target WHERE (a = ? AND (b < ? OR c != ?) AND d > ?) ORDER BY a DESC LIMIT 20
	// [test 2 5 0]
}

func newInt(i int) *int {
	return &i
}
