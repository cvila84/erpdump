package utils

import "fmt"

type Filter func(data []interface{}, index ...int) bool

type Selector func(data []interface{}, index ...int) interface{}

type Table struct {
	data            [][]interface{}
	cells           [][]interface{}
	uniques         map[int]Vector[string]
	filters         []Filter
	rowSelectors    []Selector
	columnSelectors []Selector
	valueSelectors  []Selector
}

func NewTable(data [][]interface{}) *Table {
	table := &Table{data: data}
	table.cells = make([][]interface{}, 0)
	table.filters = make([]Filter, 0)
	table.rowSelectors = make([]Selector, 0)
	table.columnSelectors = make([]Selector, 0)
	table.valueSelectors = make([]Selector, 0)
	return table
}

func (t *Table) String() string {
	return ""
}

func (t *Table) Generate() {
	for _, item := range t.data {
		for rowSelector := range t.rowSelectors {
			unique, ok := t.uniques[i]
			if !ok {
				unique = Vector[string]{ID: func(element string) string { return element }}
				t.uniques[i] = unique
			}
			unique.Add(fmt.Sprint(item[i]))
		}
	}
}

func (t *Table) Filter(filters ...Filter) *Table {
	t.filters = filters
	return t
}

func (t *Table) Rows(selectors ...Selector) *Table {
	t.rowSelectors = selectors
	return t
}

func (t *Table) Columns(selectors ...Selector) *Table {
	t.columnSelectors = selectors
	return t
}

func (t *Table) Values(selectors ...Selector) *Table {
	t.valueSelectors = selectors
	return t
}
