package table

import (
	"fmt"
	"strings"
)

type ValueAction int

const (
	Count ValueAction = iota
	Sum
)

type Filter func(elements []interface{}) []interface{}

type Compute[T string | int | float64] func(elements []T) T

type Series[T string | int | float64] struct {
	indexes []int
	compute Compute[T]
	action  ValueAction
	sorting Sorting
}

type Table[T int | float64] struct {
	data          [][]interface{}
	pivot         map[string]map[string][]T
	filters       map[int]Filter
	rowHeaders    *Headers
	columnHeaders *Headers
	rowSeries     []Series[string]
	columnSeries  []Series[string]
	valueSeries   []Series[T]
}

func NewIntTable(data [][]interface{}) *Table[int] {
	return &Table[int]{
		data:          data,
		pivot:         make(map[string]map[string][]int),
		filters:       make(map[int]Filter),
		rowHeaders:    NewRootHeaders(),
		columnHeaders: NewRootHeaders(),
		rowSeries:     make([]Series[string], 0),
		columnSeries:  make([]Series[string], 0),
		valueSeries:   make([]Series[int], 0),
	}
}

func NewFloatTable(data [][]interface{}) *Table[float64] {
	return &Table[float64]{
		data:          data,
		pivot:         make(map[string]map[string][]float64),
		filters:       make(map[int]Filter),
		rowHeaders:    NewRootHeaders(),
		columnHeaders: NewRootHeaders(),
		rowSeries:     make([]Series[string], 0),
		columnSeries:  make([]Series[string], 0),
		valueSeries:   make([]Series[float64], 0),
	}
}

func addElement[T any](elements []T, element interface{}) ([]T, error) {
	str, ok := element.(T)
	if !ok {
		return elements, fmt.Errorf("internal error during value parsing, expected %T", *new(T))
	}
	elements = append(elements, str)
	return elements, nil
}

func walk(headers *Headers, series []Series[string], record []interface{}) (string, error) {
	h := headers
	for i := 0; i < len(series); i++ {
		var value string
		var elements []string
		for j := 0; j < len(series[i].indexes); j++ {
			var err error
			elements, err = addElement(elements, record[series[i].indexes[j]])
			if err != nil {
				return "", err
			}
		}
		if series[i].compute != nil {
			value = series[i].compute(elements)
		} else {
			value = elements[0]
		}
		h = h.sortedWalk(value, series[i].sorting)
	}
	return h.label, nil
}

func (t *Table[T]) updateCell(rowLabel string, columnLabel string, valueSeries []Series[T], record []interface{}) error {
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
		rr, ok := t.pivot[rowLabel]
		if !ok {
			rr = make(map[string][]T)
			t.pivot[rowLabel] = rr
		}
		rc, ok := rr[columnLabel]
		if !ok {
			rc = make([]T, len(valueSeries))
			rr[columnLabel] = rc
		}
		switch valueSeries[i].action {
		case Count:
			rc[i]++
		case Sum:
			rc[i] += cellValue
		}
	}
	return nil
}

func (t *Table[T]) Generate() error {
	for _, record := range t.data {
		rowLabel, err := walk(t.rowHeaders, t.rowSeries, record)
		if err != nil {
			return err
		}
		columnLabel, err := walk(t.columnHeaders, t.columnSeries, record)
		if err != nil {
			return err
		}
		err = t.updateCell(rowLabel, columnLabel, t.valueSeries, record)
		if err != nil {
			return err
		}
		sumColumnLabel := columnLabel
		for i := 0; i < len(t.columnSeries)+1; i++ {
			sumRowLabel := rowLabel
			for j := 0; j < len(t.rowSeries)+1; j++ {
				if i != 0 || j != 0 {
					err = t.updateCell(sumRowLabel, sumColumnLabel, t.valueSeries, record)
					if err != nil {
						return err
					}
				}
				sumRowLabel = parentLabel(sumRowLabel)
			}
			sumColumnLabel = parentLabel(sumColumnLabel)
		}
	}
	return nil
}

func (t *Table[T]) ToCSV() string {
	var sb strings.Builder
	for _, columnLabel := range t.columnHeaders.labels(true, true) {
		if columnLabel == "" {
			_, _ = fmt.Fprint(&sb, ";Total")
		} else {
			_, _ = fmt.Fprint(&sb, ";"+columnLabel)
		}
	}
	_, _ = fmt.Fprintln(&sb)
	for _, rowLabel := range t.rowHeaders.labels(true, true) {
		if rowLabel == "" {
			_, _ = fmt.Fprint(&sb, "Total;")
		} else {
			_, _ = fmt.Fprint(&sb, rowLabel+";")
		}
		columnLabels := t.columnHeaders.labels(true, true)
		for i, columnLabel := range columnLabels {
			v, ok := t.pivot[rowLabel][columnLabel]
			if ok {
				if len(v) > 1 {
					_, _ = fmt.Fprintf(&sb, "%v", v)
				} else {
					_, _ = fmt.Fprintf(&sb, "%v", v[0])
				}
			} else {
				_, _ = fmt.Fprintf(&sb, "")
			}
			if i < len(columnLabels)-1 {
				_, _ = fmt.Fprintf(&sb, ";")
			}
		}
		_, _ = fmt.Fprintln(&sb)
	}
	return sb.String()
}

func (t *Table[T]) Filter(index int, filter Filter) *Table[T] {
	t.filters[index] = filter
	return t
}

func (t *Table[T]) Row(index int) *Table[T] {
	return t.SortedComputedRow([]int{index}, nil, AlphaSorting)
}

func (t *Table[T]) SortedRow(index int, sorting Sorting) *Table[T] {
	return t.SortedComputedRow([]int{index}, nil, sorting)
}

func (t *Table[T]) ComputedRow(indexes []int, compute Compute[string]) *Table[T] {
	return t.SortedComputedRow(indexes, compute, AlphaSorting)
}

func (t *Table[T]) SortedComputedRow(indexes []int, compute Compute[string], sorting Sorting) *Table[T] {
	t.rowSeries = append(t.rowSeries, Series[string]{indexes, compute, Count, sorting})
	return t
}

func (t *Table[T]) Column(index int) *Table[T] {
	return t.SortedComputedColumn([]int{index}, nil, AlphaSorting)
}

func (t *Table[T]) SortedColumn(index int, sorting Sorting) *Table[T] {
	return t.SortedComputedColumn([]int{index}, nil, sorting)
}

func (t *Table[T]) ComputedColumn(indexes []int, compute Compute[string]) *Table[T] {
	return t.SortedComputedColumn(indexes, compute, AlphaSorting)
}

func (t *Table[T]) SortedComputedColumn(indexes []int, compute Compute[string], sorting Sorting) *Table[T] {
	t.columnSeries = append(t.columnSeries, Series[string]{indexes, compute, Count, sorting})
	return t
}

func (t *Table[T]) Values(index int, action ValueAction) *Table[T] {
	t.valueSeries = append(t.valueSeries, Series[T]{[]int{index}, nil, action, nil})
	return t
}

func (t *Table[T]) ComputedValues(indexes []int, compute Compute[T], action ValueAction) *Table[T] {
	t.valueSeries = append(t.valueSeries, Series[T]{indexes, compute, action, nil})
	return t
}
