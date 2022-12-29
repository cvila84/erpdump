package erp

import (
	"regexp"
)

func ParseProjectID(projectName string) string {
	r, _ := regexp.Compile(".*\\((.*)\\)$")
	if r != nil {
		g := r.FindStringSubmatch(projectName)
		if len(g) > 1 {
			return g[1]
		}
	}
	return "N/A"
}
