package pivot

import "fmt"

type seriesType interface{ string | float64 }

type series[T seriesType] struct {
	indexes   []int
	name      string
	filter    Filter
	compute   Compute[T]
	operation Operation
	sort      Sort
	format    string
}

func newRCSeries(indexes []int, filter Filter, compute Compute[string], sort Sort) *series[string] {
	return &series[string]{
		indexes:   indexes,
		filter:    filter,
		compute:   compute,
		operation: none,
		sort:      sort,
	}
}

func newVSeries(name SeriesName, indexes []DataIndex, compute Compute[float64], operation Operation, format ValueFormat) *series[float64] {
	dataIndexes := make([]int, len(indexes))
	for i := 0; i < len(indexes); i++ {
		dataIndexes[i] = int(indexes[i])
	}
	return &series[float64]{
		indexes:   dataIndexes,
		name:      string(name),
		compute:   compute,
		operation: operation,
		format:    string(format),
	}
}

func (s *series[T]) NameFromHeaders(headers []interface{}) {
	if s.compute != nil {
		if len(s.indexes) > 0 && len(s.name) == 0 {
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
