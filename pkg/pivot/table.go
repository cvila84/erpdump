package pivot

import (
	"fmt"
	"strings"
)

type Cell interface {
	fmt.Stringer
	GetValue() float64
	SetValue(float64)
	AddValue(float64)
	IncValue()
}

type pivotCell struct {
	value      float64
	fmtDisplay string
}

func (p *pivotCell) String() string {
	return fmt.Sprintf(p.fmtDisplay, p.value)
}

func (p *pivotCell) GetValue() float64 {
	return p.value
}

func (p *pivotCell) SetValue(value float64) {
	p.value = value
}

func (p *pivotCell) AddValue(value float64) {
	p.value += value
}

func (p *pivotCell) IncValue() {
	p.value++
}

type Filter func(element interface{}) bool

type Sort func(elements []string) []string

type Compute[T SeriesType] func(record []interface{}) (T, error)

type Action int

const (
	Set Action = iota
	Count
	Sum
)

type SeriesType interface{ string | float64 }

type series[T SeriesType] struct {
	name    string
	indexes []int
	filter  Filter
	compute Compute[T]
	action  Action
	sort    Sort
	display string
}

func (s *series[T]) NameFromHeaders(headers []interface{}) {
	if s.compute != nil {
		if len(s.indexes) > 0 {
			s.name = fmt.Sprintf("Computed%v", s.indexes)
		}
	} else {
		if len(s.indexes) > 0 {
			if headers != nil {
				var ok bool
				s.name, ok = headers[s.indexes[0]].(string)
				if !ok {
					s.name = fmt.Sprintf("Unnamed%v", s.indexes)
				}
			} else {
				s.name = fmt.Sprintf("Unnamed%v", s.indexes)
			}
		}
	}
}

type Table struct {
	data          [][]interface{}
	dataHeaders   bool
	pivot         map[string]map[string][]Cell
	filters       map[int]Filter
	rowHeaders    *headers
	columnHeaders *headers
	rowSeries     []*series[string]
	columnSeries  []*series[string]
	valueSeries   []*series[float64]
}

func NewTable(data [][]interface{}, dataHeaders bool) *Table {
	return &Table{
		data:          data,
		dataHeaders:   dataHeaders,
		pivot:         make(map[string]map[string][]Cell),
		filters:       make(map[int]Filter),
		rowHeaders:    newRootHeaders(nil),
		columnHeaders: newRootHeaders(nil),
		rowSeries:     make([]*series[string], 0),
		columnSeries:  make([]*series[string], 0),
		valueSeries:   make([]*series[float64], 0),
	}
}

func (t *Table) updateCellByCompute(rowLabel string, columnLabel string, record []interface{}, onlyCompute bool) error {
	for is, serie := range t.valueSeries {
		if (serie.compute != nil && onlyCompute) || (serie.compute == nil && !onlyCompute) {
			rr, ok := t.pivot[rowLabel]
			if !ok {
				rr = make(map[string][]Cell)
				t.pivot[rowLabel] = rr
			}
			rc, ok := rr[columnLabel]
			if !ok {
				rc = make([]Cell, len(t.valueSeries))
				rr[columnLabel] = rc
			}
			if rc[is] == nil {
				rc[is] = &pivotCell{
					fmtDisplay: serie.display,
				}
			}
			switch serie.action {
			case Count:
				rc[is].IncValue()
			case Sum:
				value, err := computeFloat(*serie, record)
				if err != nil {
					return fmt.Errorf("while updating cell [%q,%q] with record %v: %w", rowLabel, columnLabel, record, err)
				}
				rc[is].AddValue(value)
			case Set:
				value, err := computeFloat(*serie, record)
				if err != nil {
					return fmt.Errorf("while updating cell [%q,%q] with record %v: %w", rowLabel, columnLabel, record, err)
				}
				rc[is].SetValue(value)
			}
		}
	}
	return nil
}

func (t *Table) updateCell(rowLabel string, columnLabel string, record []interface{}) error {
	err := t.updateCellByCompute(rowLabel, columnLabel, record, false)
	if err != nil {
		return err
	}
	return t.updateCellByCompute(rowLabel, columnLabel, record, true)
}

func (t *Table) updateCrossCells(rowLabel string, columnLabel string, record []interface{}) error {
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

func (t *Table) Generate() error {
	var headerSeries []*series[string]
	var headerLabels []interface{}
	if t.dataHeaders {
		headerLabels = t.data[0]
	}
	headerSeries = append(headerSeries, t.rowSeries...)
	headerSeries = append(headerSeries, t.columnSeries...)
	for _, serie := range headerSeries {
		serie.NameFromHeaders(headerLabels)
	}
	filteredData, err := filter(t.filters, headerSeries, t.data, t.dataHeaders)
	if err != nil {
		return err
	}
	for _, serie := range t.valueSeries {
		serie.NameFromHeaders(headerLabels)
	}
	for _, record := range filteredData {
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
			return fmt.Errorf("empty column labels are not supported")
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

// ToCSV
// TODO manage multi-values through virtual column
func (t *Table) ToCSV() string {
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
						sbv.WriteString(v[j].String())
						if j != len(v)-1 {
							sbv.WriteString(", ")
						}
					}
					sbv.WriteString(" ]")
					_, _ = fmt.Fprintf(&sb, "%s", sbv.String())
				} else {
					_, _ = fmt.Fprintf(&sb, "%v", v[0].String())
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

func (t *Table) Filter(index int, filter Filter) *Table {
	t.filters[index] = filter
	return t
}

func (t *Table) StandardRow(index int) *Table {
	return t.Row([]int{index}, nil, nil, nil)
}

func (t *Table) Row(indexes []int, filter Filter, compute Compute[string], sort Sort) *Table {
	t.rowSeries = append(t.rowSeries, &series[string]{
		indexes: indexes,
		compute: compute,
		filter:  filter,
		sort:    sort,
		action:  Set,
	})
	return t
}

func (t *Table) StandardColumn(index int) *Table {
	return t.Column([]int{index}, nil, nil, nil)
}

func (t *Table) Column(indexes []int, filter Filter, compute Compute[string], sort Sort) *Table {
	t.columnSeries = append(t.columnSeries, &series[string]{
		indexes: indexes,
		filter:  filter,
		compute: compute,
		sort:    sort,
		action:  Set,
	})
	return t
}

func (t *Table) StandardValues(index int, action Action) *Table {
	t.valueSeries = append(t.valueSeries, &series[float64]{
		indexes: []int{index},
		sort:    nil,
		compute: nil,
		action:  action,
		display: Digits(0),
	})
	return t
}

func (t *Table) Values(indexes []int, compute Compute[float64], action Action, display string) *Table {
	t.valueSeries = append(t.valueSeries, &series[float64]{
		indexes: indexes,
		sort:    nil,
		compute: compute,
		action:  action,
		display: display,
	})
	return t
}
