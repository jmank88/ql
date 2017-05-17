// Package ql contains functions for adapting generic JSON database queries to different backend databases
package ql

// Query is a generic, JSON-serializable query.
type Query struct {
	// Target of the query.
	Target string `json:"target"`
	// Optional names of fields to fetch.
	Fields []string `json:"fields,omitempty"`
	// Optional filter.
	Filter *Filter `json:"filter,omitempty"`
	// Optional sort order.
	Sort []Sort `json:"sort,omitempty"`
	// Optional maximum records to fetch.
	Max *int `json:"max,omitempty"`
}

// Asc creates an ascending Sort.
func Asc(field string) Sort {
	return Sort{Field: field}
}

// Desc creates a descending Sort.
func Desc(field string) Sort {
	return Sort{Field: field, Descending: true}
}

// Sort defines a sort ordering for a Field.
type Sort struct {
	Field      string `json:"field"`
	Descending bool   `json:"descending"`
}

// And creates an 'AND' Composite Filter.
func And(filters ...*Filter) *Filter {
	c := &Filter{
		Composite: Composite{
			Operator: "AND",
			Filters:  filters,
		},
	}
	return c
}

// Or creates an 'OR' Composite Filter.
func Or(filters ...*Filter) *Filter {
	c := &Filter{
		Composite: Composite{
			Operator: "OR",
			Filters:  filters,
		},
	}
	return c
}

// Compare creates a Comparison Filter.
func Compare(field, comparator string, value interface{}) *Filter {
	return &Filter{
		Comparison: Comparison{
			Field:      field,
			Comparator: comparator,
			Value:      value,
		},
	}
}

// Filter is a union of filter types. Only one embedded struct should have set fields.
type Filter struct {
	Comparison
	Composite
}

// A Comparison filter compares a field to a value.
type Comparison struct {
	// One of: =, !=, <, >, <=, >=.
	Comparator string      `json:"comparator,omitempty"`
	Field      string      `json:"field,omitempty"`
	Value      interface{} `json:"value,omitempty"`
}

// An Composite filter logically combines two or more Filters.
type Composite struct {
	// One of (case-insensitive): AND, OR.
	Operator string    `json:"operator,omitempty"`
	Filters  []*Filter `json:"filters,omitempty"`
}
