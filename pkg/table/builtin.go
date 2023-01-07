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

var Regroup = func(list []string, memberLabel string, nonMemberLabel string) Compute[string] {
	return func(elements []string) string {
		for _, e := range list {
			if elements[0] == e {
				return memberLabel
			}
		}
		return nonMemberLabel
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
