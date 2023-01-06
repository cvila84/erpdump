package table

import (
	"fmt"
	"sort"
)

type ValueAction int

const (
	Count ValueAction = iota
	Sum
)

type Filter func(elements []interface{}) []interface{}

type Compute[T string | int | float64] func(elements []T) T

type Sorting func(elements []string) []string

var AlphaSorting Sorting = func(elements []string) []string {
	sort.Strings(elements)
	return elements
}

type Series[T string | int | float64] struct {
	indexes []int
	action  ValueAction
	compute Compute[T]
}

type Table[T int | float64] struct {
	data          [][]interface{}
	oldcells      map[string]interface{}
	pivot         map[string]map[string]interface{}
	rowHeaders    Headers
	columnHeaders Headers
	filters       map[int]Filter
	rowSeries     []Series[string]
	columnSeries  []Series[string]
	valueSeries   []Series[T]
}

func NewIntTable(data [][]interface{}) *Table[int] {
	return &Table[int]{
		data:         data,
		oldcells:     make(map[string]interface{}),
		filters:      make(map[int]Filter),
		rowSeries:    make([]Series[string], 0),
		columnSeries: make([]Series[string], 0),
		valueSeries:  make([]Series[int], 0),
	}
}

func NewFloatTable(data [][]interface{}) *Table[float64] {
	return &Table[float64]{
		data:         data,
		oldcells:     make(map[string]interface{}),
		filters:      make(map[int]Filter),
		rowSeries:    make([]Series[string], 0),
		columnSeries: make([]Series[string], 0),
		valueSeries:  make([]Series[float64], 0),
	}
}

func addElement[T any](elements []T, element interface{}) ([]T, error) {
	str, ok := element.(T)
	if !ok {
		return elements, fmt.Errorf("internal error during row value parsing, expected %T", *new(T))
	}
	elements = append(elements, str)
	return elements, nil
}

func walk(headers *Headers, series []Series[string], record []interface{}) {

}

func getRow(rows *map[string]interface{}, rowSeries []Series[string], record []interface{}) (*map[string]interface{}, error) {
	m := rows
	for i := 0; i < len(rowSeries); i++ {
		var rowValue string
		var elements []string
		for j := 0; j < len(rowSeries[i].indexes); j++ {
			var err error
			elements, err = addElement(elements, record[rowSeries[i].indexes[j]])
			if err != nil {
				return nil, err
			}
		}
		if rowSeries[i].compute != nil {
			rowValue = rowSeries[i].compute(elements)
		} else {
			rowValue = elements[0]
		}
		e, ok := (*m)[rowValue]
		if !ok {
			e = make(map[string]interface{})
			(*m)[rowValue] = e
		}
		nm, ok := e.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("internal error during rows parsing, expected map")
		}
		m = &nm
	}
	return m, nil
}

func getCell[T int | float64](row *map[string]interface{}, columnSeries []Series[string], cellSize int, record []interface{}) (*[]T, error) {
	m := row
	for i := 0; i < len(columnSeries); i++ {
		var columnValue string
		var elements []string
		for j := 0; j < len(columnSeries[i].indexes); j++ {
			var err error
			elements, err = addElement(elements, record[columnSeries[i].indexes[j]])
			if err != nil {
				return nil, err
			}
		}
		if columnSeries[i].compute != nil {
			columnValue = columnSeries[i].compute(elements)
		} else {
			columnValue = elements[0]
		}
		e, ok := (*m)[columnValue]
		if !ok {
			if i != len(columnSeries)-1 {
				e = make(map[string]interface{})
			} else {
				e = make([]T, cellSize)
			}
			(*m)[columnValue] = e
		}
		if i != len(columnSeries)-1 {
			var nm map[string]interface{}
			nm, ok = e.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("internal error during column parsing, expected map")
			}
			m = &nm
		} else {
			var c []T
			c, ok = e.([]T)
			if !ok {
				return nil, fmt.Errorf("internal error during cell parsing, expected slice")
			}
			return &c, nil
		}
	}
	return nil, fmt.Errorf("internal error during column parsing, no columns defined")
}

func updateCell[T int | float64](cell *[]T, valueSeries []Series[T], record []interface{}) error {
	for i := 0; i < len(valueSeries); i++ {
		var cellValue T
		var elements []T
		for j := 0; j < len(valueSeries[i].indexes); j++ {
			var err error
			elements, err = addElement(elements, record[valueSeries[i].indexes[j]])
			if err != nil {
				return err
			}
		}
		if valueSeries[i].compute != nil {
			cellValue = valueSeries[i].compute(elements)
		} else {
			cellValue = elements[0]
		}
		switch valueSeries[i].action {
		case Count:
			(*cell)[i]++
		case Sum:
			(*cell)[i] += cellValue
		}
	}
	return nil
}

func (t *Table[T]) Generate() error {
	for _, record := range t.data {
		row, err := getRow(&t.oldcells, t.rowSeries, record)
		if err != nil {
			return err
		}
		cell, err := getCell[T](row, t.columnSeries, len(t.valueSeries), record)
		if err != nil {
			return err
		}
		err = updateCell[T](cell, t.valueSeries, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Table[T]) ToCSV() (string, error) {
	//e := t.oldcells
	//for i := 0; i < len(t.rowSeries); i++ {
	//	for k, v := range e {
	//	}
	//}
	return "", nil
}

func (t *Table[T]) Filter(index int, filter Filter) *Table[T] {
	t.filters[index] = filter
	return t
}

func (t *Table[T]) Row(index int, sorting Sorting) *Table[T] {
	t.rowSeries = append(t.rowSeries, Series[string]{[]int{index}, Count, nil})
	return t
}

func (t *Table[T]) ComputedRow(indexes []int, compute Compute[string], sorting Sorting) *Table[T] {
	t.rowSeries = append(t.rowSeries, Series[string]{indexes, Count, compute})
	return t
}

func (t *Table[T]) Column(index int, sorting Sorting) *Table[T] {
	t.columnSeries = append(t.columnSeries, Series[string]{[]int{index}, Count, nil})
	return t
}

func (t *Table[T]) ComputedColumn(indexes []int, compute Compute[string], sorting Sorting) *Table[T] {
	t.columnSeries = append(t.columnSeries, Series[string]{indexes, Count, compute})
	return t
}

func (t *Table[T]) Values(index int, action ValueAction) *Table[T] {
	t.valueSeries = append(t.valueSeries, Series[T]{[]int{index}, action, nil})
	return t
}

func (t *Table[T]) ComputedValues(indexes []int, compute Compute[T], action ValueAction) *Table[T] {
	t.valueSeries = append(t.valueSeries, Series[T]{indexes, action, compute})
	return t
}
