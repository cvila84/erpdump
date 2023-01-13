package pivot

import (
	"fmt"
)

func valueString(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	}
	return fmt.Sprintf("%v", v)
}

func compute[T headerTypes | valueTypes](t []T, serie series[T], record []interface{}) (T, error) {
	var value T
	var previousElements []T
	var computeElements []interface{}
	for _, i := range serie.indexes {
		previousElements = append(previousElements, t[i])
		computeElements = append(computeElements, record[i])
	}
	if serie.compute != nil {
		var err error
		value, err = serie.compute(computeElements)
		if err != nil {
			return *new(T), fmt.Errorf("while computing for %v: %w", computeElements, err)
		}
	} else {
		var ok bool
		value, ok = computeElements[0].(T)
		if !ok {
			return *new(T), fmt.Errorf("invalid type %T for element %s", computeElements[0], computeElements[0])
		}
	}
	return value, nil
}

func filter(filters map[int]Filter, series []series[string], records [][]interface{}) ([][]interface{}, error) {
	var filteredRecords [][]interface{}
	for _, record := range records {
		keep := true
		for i, f := range filters {
			if !f(record[i]) {
				keep = false
			}
		}
		for _, serie := range series {
			value, err := compute(serie, record)
			if err != nil {
				return nil, err
			}
			if serie.filter != nil && !serie.filter(value) {
				keep = false
			}
		}
		if keep {
			filteredRecords = append(filteredRecords, record)
		}
	}
	return filteredRecords, nil
}

func walk(headers *headers, series []series[string], record []interface{}) (string, error) {
	h := headers
	for _, serie := range series {
		value, err := compute(serie, record)
		if err != nil {
			return "", err
		}
		h = h.sort(serie.sort).walk(value)
	}
	return h.label, nil
}
