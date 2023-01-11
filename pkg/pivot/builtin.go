package pivot

import (
	"fmt"
	"sort"
	"strings"
)

var AlphaSort Sort = func(elements []string) []string {
	less := func(i, j int) bool {
		return strings.ToLower(elements[i]) < strings.ToLower(elements[j])
	}
	sort.SliceStable(elements, less)
	return elements
}

var ReverseAlphaSort Sort = func(elements []string) []string {
	less := func(i, j int) bool {
		return strings.ToLower(elements[i]) > strings.ToLower(elements[j])
	}
	sort.SliceStable(elements, less)
	return elements
}

var Group = func(groups [][]string, groupLabels []string, noneLabel string) Compute[string] {
	return func(elements []interface{}) (string, error) {
		for i, group := range groups {
			for _, groupElement := range group {
				e, ok := elements[0].(string)
				if !ok {
					return "", fmt.Errorf("invalid type %T for element %s", elements[0], elements[0])
				}
				if e == groupElement {
					return groupLabels[i], nil
				}
			}
		}
		return noneLabel, nil
	}
}

var SumFloats Compute[float64] = func(elements []interface{}) (float64, error) {
	var result float64
	for _, element := range elements {
		f, ok := element.(float64)
		if !ok {
			return 0, fmt.Errorf("invalid type %T for element %s", element, element)
		}
		result += f
	}
	return result, nil
}

var PartialSumFloats = func(sumGroup, groupSize int) Compute[float64] {
	return func(elements []interface{}) (float64, error) {
		var result float64
		for i, element := range elements {
			e, ok := element.(float64)
			if !ok {
				return 0, fmt.Errorf("invalid type %T for element %s", element, element)
			}
			if i >= groupSize*(sumGroup-1) && i < groupSize*sumGroup {
				result += e
			}
		}
		return result, nil
	}
}

var In = func(list []string) Filter {
	return func(element interface{}) bool {
		for _, e := range list {
			if element == e {
				return true
			}
		}
		return false
	}
}
