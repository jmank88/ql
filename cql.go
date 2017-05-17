package ql

import (
	"bytes"
	"fmt"
	"strings"
)

type cqlBuilder struct {
	bytes.Buffer
	args []interface{}
}

// CQL returns a CQL query and arguments.
func CQL(q *Query) (string, []interface{}) {
	b := &cqlBuilder{}
	fmt.Fprintf(b, "SELECT %s FROM %s", strings.Join(q.Fields, ", "), q.Target)
	if q.Filter != nil {
		fmt.Fprint(b, " WHERE ")
		b.filter(q.Filter)
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

func (b *cqlBuilder) composite(c *Composite) {
	fmt.Fprint(b, "(")
	if len(c.Filters) > 0 {

		b.filter(c.Filters[0])
		for _, f := range c.Filters[1:] {
			fmt.Fprintf(b, " %s ", c.Operator)
			b.filter(f)
		}
	}
	fmt.Fprint(b, ")")
}

func (b *cqlBuilder) filter(f *Filter) {
	if len(f.Filters) > 0 {
		b.composite(&f.Composite)
	} else {
		b.comparison(&f.Comparison)
	}
}

func (b *cqlBuilder) comparison(c *Comparison) {
	b.args = append(b.args, c.Value)
	fmt.Fprintf(b, "%s %s ?", c.Field, c.Comparator)
}

func (b *cqlBuilder) sort(s Sort) {
	fmt.Fprint(b, s.Field)
	if s.Descending {
		fmt.Fprint(b, " DESC")
	}
}
