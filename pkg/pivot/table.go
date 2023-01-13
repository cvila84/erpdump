package pivot

import (
	"fmt"
	"strings"
)

type headerTypes interface{ string }
type valueTypes interface{ int | float64 }

type Filter func(element interface{}) bool

type Sort func(elements []string) []string

type Compute[T headerTypes | valueTypes] func(record []interface{}) (T, error)

type Action int

const (
	None Action = iota
	Count
	Sum
)

type series[T headerTypes | valueTypes] struct {
	indexes []int
	filter  Filter
	compute Compute[T]
	action  Action
	sort    Sort
}

type Table[T valueTypes] struct {
	data          [][]interface{}
	filteredData  [][]interface{}
	pivot         map[string]map[string][]T
	filters       map[int]Filter
	rowHeaders    *headers
	columnHeaders *headers
	rowSeries     []series[string]
	columnSeries  []series[string]
	valueSeries   []series[T]
}

func NewTable[T valueTypes](data [][]interface{}) *Table[T] {
	return &Table[T]{
		data:          data,
		filteredData:  nil,
		pivot:         make(map[string]map[string][]T),
		filters:       make(map[int]Filter),
		rowHeaders:    newRootHeaders(nil),
		columnHeaders: newRootHeaders(nil),
		rowSeries:     make([]series[string], 0),
		columnSeries:  make([]series[string], 0),
		valueSeries:   make([]series[T], 0),
	}
}

func (t *Table[T]) updateCell(rowLabel string, columnLabel string, record []interface{}) error {
	for is, serie := range t.valueSeries {
		rr, ok := t.pivot[rowLabel]
		if !ok {
			rr = make(map[string][]T)
			t.pivot[rowLabel] = rr
		}
		rc, ok := rr[columnLabel]
		if !ok {
			rc = make([]T, len(t.valueSeries))
			rr[columnLabel] = rc
		}
		value, err := compute(rc, serie, record)
		if err != nil {
			return err
		}
		rc[is] = value
	}
	return nil
}

func (t *Table[T]) updateCrossCells(rowLabel string, columnLabel string, record []interface{}) error {
	sumColumnLabel := columnLabel
	for i := 0; i < len(t.columnSeries)+1; i++ {
		sumRowLabel := rowLabel
		for j := 0; j < len(t.rowSeries)+1; j++ {
			if i != 0 || j != 0 {
				err := t.updateCell(sumRowLabel, sumColumnLabel, record)
				if err != nil {
					return err
				}
			}
			sumRowLabel = parentHeaderLabel(sumRowLabel)
		}
		sumColumnLabel = parentHeaderLabel(sumColumnLabel)
	}
	return nil
}

func (t *Table[T]) Generate() error {
	var allSeries []series[string]
	var err error
	allSeries = append(allSeries, t.rowSeries...)
	allSeries = append(allSeries, t.columnSeries...)
	t.filteredData, err = filter(t.filters, allSeries, t.data)
	if err != nil {
		return err
	}
	for _, record := range t.filteredData {
		var rowLabel string
		var columnLabel string
		rowLabel, err = walk(t.rowHeaders, t.rowSeries, record)
		if err != nil {
			return err
		}
		if len(rowLabel) == 0 {
			return fmt.Errorf("empty row labels are not supported")
		}
		columnLabel, err = walk(t.columnHeaders, t.columnSeries, record)
		if err != nil {
			return err
		}
		if len(columnLabel) == 0 {
			return fmt.Errorf("empty row labels are not supported")
		}
		err = t.updateCell(rowLabel, columnLabel, record)
		if err != nil {
			return err
		}
		err = t.updateCrossCells(rowLabel, columnLabel, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Table[T]) ToCSV() string {
	columnLabels := t.columnHeaders.labels(true, true)
	rowLabels := t.rowHeaders.labels(true, true)
	var sb strings.Builder
	for _, columnLabel := range columnLabels {
		if columnLabel == "" {
			_, _ = fmt.Fprint(&sb, ";Total")
		} else {
			_, _ = fmt.Fprint(&sb, ";"+columnLabel)
		}
	}
	_, _ = fmt.Fprintln(&sb)
	for _, rowLabel := range rowLabels {
		if rowLabel == "" {
			_, _ = fmt.Fprint(&sb, "Total;")
		} else {
			_, _ = fmt.Fprint(&sb, rowLabel+";")
		}
		for i, columnLabel := range columnLabels {
			v, ok := t.pivot[rowLabel][columnLabel]
			if ok {
				if len(v) > 1 {
					var sbv strings.Builder
					sbv.WriteString("[ ")
					for j := 0; j < len(v); j++ {
						sbv.WriteString(valueString(v[j]))
						if j != len(v)-1 {
							sbv.WriteString(", ")
						}
					}
					sbv.WriteString(" ]")
					_, _ = fmt.Fprintf(&sb, "%s", sbv.String())
				} else {
					_, _ = fmt.Fprintf(&sb, "%v", valueString(v[0]))
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

func (t *Table[T]) StandardRow(index int) *Table[T] {
	return t.Row([]int{index}, nil, nil, nil)
}

func (t *Table[T]) Row(indexes []int, compute Compute[string], filter Filter, sort Sort) *Table[T] {
	t.rowSeries = append(t.rowSeries, series[string]{
		indexes: indexes,
		filter:  filter,
		sort:    sort,
		compute: compute,
		action:  None,
	})
	return t
}

func (t *Table[T]) StandardColumn(index int) *Table[T] {
	return t.Column([]int{index}, nil, nil, nil)
}

func (t *Table[T]) Column(indexes []int, filter Filter, compute Compute[string], sort Sort) *Table[T] {
	t.columnSeries = append(t.columnSeries, series[string]{
		indexes: indexes,
		filter:  filter,
		sort:    sort,
		compute: compute,
		action:  None,
	})
	return t
}

func (t *Table[T]) StandardValues(index int, action Action) *Table[T] {
	return t.Values([]int{index}, nil, nil, action)
}

func (t *Table[T]) Values(indexes []int, filter Filter, compute Compute[T], action Action) *Table[T] {
	t.valueSeries = append(t.valueSeries, series[T]{
		indexes: indexes,
		filter:  filter,
		sort:    nil,
		compute: compute,
		action:  action,
	})
	return t
}
