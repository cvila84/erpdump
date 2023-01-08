package table

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

var Group1 = func(group1 []string, group1Label string, nowhereLabel string) Compute[string] {
	return func(elements []string) string {
		for _, e := range group1 {
			if elements[0] == e {
				return group1Label
			}
		}
		return nowhereLabel
	}
}

var Group2 = func(group1 []string, group2 []string, group1Label string, group2Label string, nowhereLabel string) Compute[string] {
	return func(elements []string) string {
		for _, e := range group1 {
			if elements[0] == e {
				return group1Label
			}
		}
		for _, e := range group2 {
			if elements[0] == e {
				return group2Label
			}
		}
		return nowhereLabel
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
