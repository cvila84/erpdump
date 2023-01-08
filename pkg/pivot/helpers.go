package pivot

import (
	"fmt"
	"strings"
)

func valueString(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	}
	return "#"
}

func parentLabel(label string) string {
	if label == "" {
		return ""
	}
	idx := strings.LastIndex(label, "/")
	if idx <= 0 {
		return ""
	}
	return label[0:idx]
}

func add[T headerTypes | valueTypes](elements []T, element interface{}) ([]T, error) {
	str, ok := element.(T)
	if !ok {
		return elements, fmt.Errorf("internal error during value parsing, expected %T", *new(T))
	}
	elements = append(elements, str)
	return elements, nil
}

func compute[T headerTypes | valueTypes](serie series[T], record []interface{}) (T, error) {
	var value T
	var elements []T
	for _, i := range serie.indexes {
		var err error
		elements, err = add(elements, record[i])
		if err != nil {
			return *new(T), err
		}
	}
	if serie.compute != nil {
		value = serie.compute(elements)
	} else {
		value = elements[0]
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
