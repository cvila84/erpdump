package utils

var InList = func(list []string, trueValue string, falseValue string) func(elements []string) string {
	return func(elements []string) string {
		for _, e := range list {
			if elements[0] == e {
				return trueValue
			}
		}
		return falseValue
	}
}
