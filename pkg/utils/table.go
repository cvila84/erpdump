package utils

import "fmt"

type Filter func(data interface{}) bool

type Table struct {
	data    [][]interface{}
	rows    [][]interface{}
	uniques map[int]Vector[string]
}

func NewTable(data [][]interface{}) *Table {
	table := &Table{data: data}
	table.rows = make([][]interface{}, 0)
	return table
}

func (t *Table) String() string {
	return ""
}

func (t *Table) Filter(index int, filter Filter) *Table {
	return t
}

func (t *Table) Rows(indexes ...int) *Table {
	for _, item := range t.data {
		for i := range indexes {
			unique, ok := t.uniques[i]
			if !ok {
				unique = Vector[string]{ID: func(element string) string { return element }}
				t.uniques[i] = unique
			}
			unique.Add(fmt.Sprint(item[i]))
		}
	}
	for i := range indexes {
		for _, item := range t.data {

		}
	}
	return t
}

func (t *Table) Columns() *Table {
	return t
}

func (t *Table) Values() *Table {
	return t
}
