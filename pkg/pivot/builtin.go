package pivot

import "sort"

var AlphaSort Sort = func(elements []string) []string {
	sort.Strings(elements)
	return elements
}

var ReverseAlphaSort Sort = func(elements []string) []string {
	sort.Strings(elements)
	for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
		elements[i], elements[j] = elements[j], elements[i]
	}
	return elements
}

var Group = func(groups [][]string, groupLabels []string, noneLabel string) Compute[string] {
	return func(elements []string) string {
		for i, group := range groups {
			for _, e := range group {
				if elements[0] == e {
					return groupLabels[i]
				}
			}
		}
		return noneLabel
	}
}

var YearlyHours Compute[float64] = func(elements []float64) float64 {
	var result float64
	for _, element := range elements {
		result += element
	}
	return result
}

var QuaterlyHours = func(quarter int) Compute[float64] {
	return func(elements []float64) float64 {
		var result float64
		for i, element := range elements {
			if i >= 3*(quarter-1) && i < 3*quarter {
				result += element
			}
		}
		return result
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
