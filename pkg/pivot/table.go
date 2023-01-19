package pivot

import (
	"fmt"
	"strings"
)

type cell interface {
	fmt.Stringer
	SetValue(float64)
	AddValue(float64)
}

type pivotCell struct {
	value      float64
	fmtDisplay string
}

func (p *pivotCell) String() string {
	return fmt.Sprintf(p.fmtDisplay, p.value)
}

func (p *pivotCell) SetValue(value float64) {
	p.value = value
}

func (p *pivotCell) AddValue(value float64) {
	p.value += value
}

type Filter func(element interface{}) bool

type Sort func(elements []string) []string

type Compute[T SeriesType] func(record []interface{}) (T, error)

type Action int

const (
	none Action = iota
	Count
	Sum
	set
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
	pivot         map[string]map[string][]cell
	filters       map[int]Filter
	rowHeaders    *headers
	columnHeaders *headers
	// TODO populate this variable when several values are requested and use it for Generate
	valueHeaders *headers
	rowSeries    []*series[string]
	columnSeries []*series[string]
	valueSeries  []*series[float64]
	indexes      map[int]int
	err          error
}

func NewTable(data [][]interface{}, dataHeaders bool) *Table {
	var err error
	if data == nil || (len(data) == 0 && !dataHeaders) || (len(data) <= 1 && dataHeaders) {
		err = fmt.Errorf("no input data")
	} else if len(data[0]) == 0 {
		err = fmt.Errorf("no input data")
	} else {
		length := -1
		for _, record := range data {
			if length < 0 {
				length = len(record)
			} else if len(record) != length {
				err = fmt.Errorf("input data has variable records size")
			}
		}
	}
	return &Table{
		data:          data,
		dataHeaders:   dataHeaders,
		pivot:         make(map[string]map[string][]cell),
		filters:       make(map[int]Filter),
		rowHeaders:    newRootHeaders(nil),
		columnHeaders: newRootHeaders(nil),
		rowSeries:     make([]*series[string], 0),
		columnSeries:  make([]*series[string], 0),
		valueSeries:   make([]*series[float64], 0),
		indexes:       make(map[int]int, 0),
		err:           err,
	}
}

func (t *Table) updateCellByCompute(rowLabel string, columnLabel string, record []interface{}, onlyCompute bool) error {
	for is, serie := range t.valueSeries {
		if (serie.compute != nil && onlyCompute) || (serie.compute == nil && !onlyCompute) {
			rr, ok := t.pivot[rowLabel]
			if !ok {
				rr = make(map[string][]cell)
				t.pivot[rowLabel] = rr
			}
			rc, ok := rr[columnLabel]
			if !ok {
				rc = make([]cell, len(t.valueSeries))
				rr[columnLabel] = rc
			}
			if rc[is] == nil {
				rc[is] = &pivotCell{
					fmtDisplay: serie.display,
				}
			}
			switch serie.action {
			case Count:
				rc[is].AddValue(1)
			case Sum:
				value, err := computeFloatWithRecord(*serie, record)
				if err != nil {
					return fmt.Errorf("while updating cell [%q,%q] with record %v: %w", rowLabel, columnLabel, record, err)
				}
				rc[is].AddValue(value)
			case set:
				value, err := computeFloatWithCell(*serie, rc)
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

func (t *Table) registerRow(indexes []int, filter Filter, compute Compute[string], sort Sort) error {
	if len(indexes) == 0 {
		return fmt.Errorf("invalid row definition, no indexes given")
	}
	if compute == nil && len(indexes) != 1 {
		return fmt.Errorf("invalid row definition, several indexes with no compute given")
	}
	if compute == nil {
		_, ok := t.indexes[indexes[0]]
		if ok {
			return fmt.Errorf("invalid row definition, index already used")
		}
		t.indexes[indexes[0]] = len(t.rowSeries)
	}
	t.rowSeries = append(t.rowSeries, &series[string]{
		indexes: indexes,
		filter:  filter,
		compute: compute,
		action:  none,
		sort:    sort,
	})
	return nil
}

func (t *Table) registerColumn(indexes []int, filter Filter, compute Compute[string], sort Sort) error {
	if len(indexes) == 0 {
		return fmt.Errorf("invalid column definition, no indexes given")
	}
	if compute == nil && len(indexes) != 1 {
		return fmt.Errorf("invalid column definition, several indexes with no compute given")
	}
	if compute == nil {
		_, ok := t.indexes[indexes[0]]
		if ok {
			return fmt.Errorf("invalid column definition, index already used")
		}
		t.indexes[indexes[0]] = len(t.columnSeries)
	}
	t.columnSeries = append(t.columnSeries, &series[string]{
		indexes: indexes,
		filter:  filter,
		compute: compute,
		action:  none,
		sort:    sort,
	})
	return nil
}

func (t *Table) registerValue(indexes []int, compute Compute[float64], action Action, display string) error {
	if len(indexes) == 0 {
		return fmt.Errorf("invalid value definition, no indexes given")
	}
	if compute == nil && len(indexes) != 1 {
		return fmt.Errorf("invalid value definition, several indexes with no compute given")
	}
	if compute == nil {
		_, ok := t.indexes[indexes[0]]
		if ok {
			return fmt.Errorf("invalid value definition, index already used")
		}
		t.indexes[indexes[0]] = len(t.valueSeries)
	} else {
		_, ok := t.indexes[indexes[0]]
		if !ok {
			t.indexes[indexes[0]] = len(t.valueSeries)
		}
	}
	t.valueSeries = append(t.valueSeries, &series[float64]{
		indexes: indexes,
		compute: compute,
		action:  action,
		display: display,
	})
	return nil
}

func (t *Table) Generate() error {
	if t.err != nil {
		return t.err
	}
	if len(t.rowSeries) == 0 {
		return fmt.Errorf("no rows defined")
	}
	if len(t.columnSeries) == 0 {
		return fmt.Errorf("no columns defined")
	}
	if len(t.valueSeries) == 0 {
		return fmt.Errorf("no values defined")
	}
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

func (t *Table) Row(index int) *Table {
	return t.ComputedRow([]int{index}, nil, nil, nil)
}

func (t *Table) ComputedRow(indexes []int, filter Filter, compute Compute[string], sort Sort) *Table {
	err := t.registerRow(indexes, filter, compute, sort)
	if t.err == nil {
		t.err = err
	}
	return t
}

func (t *Table) Column(index int) *Table {
	return t.ComputedColumn([]int{index}, nil, nil, nil)
}

func (t *Table) ComputedColumn(indexes []int, filter Filter, compute Compute[string], sort Sort) *Table {
	err := t.registerColumn(indexes, filter, compute, sort)
	if t.err == nil {
		t.err = err
	}
	return t
}

func (t *Table) Values(index int, action Action, display string) *Table {
	err := t.registerValue([]int{index}, nil, action, display)
	if t.err == nil {
		t.err = err
	}
	return t
}

// TODO give a name to computed values so we can display it in Generate when we have several values
func (t *Table) ComputedValues(indexes []int, compute Compute[float64], display string) *Table {
	err := t.registerValue(indexes, compute, set, display)
	if t.err == nil {
		t.err = err
	}
	return t
}
