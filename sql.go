package ql

import (
	"bytes"
	"fmt"
	"strings"
)

// MYSQL returns a MYSQL compatible SQL query and arguments.
func MYSQL(q *Query) (string, []interface{}) {
	b := &sqlBuilder{sprintArg: questionArgs}
	return b.query(q)
}

// POSTGRESQL returns a POSTGRESQL compatible SQL query and arguments.
func POSTGRESQL(q *Query) (string, []interface{}) {
	b := &sqlBuilder{sprintArg: numberedArg}
	return b.query(q)
}

type sqlBuilder struct {
	bytes.Buffer
	args      []interface{}
	sprintArg func(int) string
}

func (b *sqlBuilder) query(q *Query) (string, []interface{}) {
	fmt.Fprintf(b, "SELECT %s FROM %s", strings.Join(q.Fields, ", "), q.Target)
	if q.Filter != nil {
		fmt.Fprint(b, " WHERE ")
		b.constraint(q.Filter)
	}
	if len(q.Sort) > 0 {
		fmt.Fprint(b, " ORDER BY ")
		b.sort(q.Sort[0])
		for _, s := range q.Sort[1:] {
			fmt.Fprint(b, ", ")
			b.sort(s)
		}
	}
	if q.Max != nil {
		fmt.Fprintf(b, " LIMIT %d", *q.Max)
	}
	return b.String(), b.args
}

func numberedArg(n int) string {
	return fmt.Sprintf("$%d", n)
}

func questionArgs(n int) string {
	return "?"
}

func (b *sqlBuilder) operation(o *Composite) {
	fmt.Fprint(b, "(")
	if len(o.Filters) > 0 {
		b.constraint(o.Filters[0])
		for _, c := range o.Filters[1:] {
			fmt.Fprintf(b, " %s ", o.Operator)
			b.constraint(c)
		}
	}
	fmt.Fprint(b, ")")
}

func (b *sqlBuilder) constraint(c *Filter) {
	if len(c.Filters) > 0 {
		b.operation(&c.Composite)
	} else {
		b.comparison(&c.Comparison)
	}
}

func (b *sqlBuilder) comparison(c *Comparison) {
	b.args = append(b.args, c.Value)
	fmt.Fprintf(b, "%s %s %s", c.Field, c.Comparator, b.sprintArg(len(b.args)))
}

func (b *sqlBuilder) sort(c Sort) {
	fmt.Fprint(b, c.Field)
	if c.Descending {
		fmt.Fprint(b, " DESC")
	}
}
