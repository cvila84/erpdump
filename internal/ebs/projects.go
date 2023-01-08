package ebs

import (
	"regexp"
)

var otaProjects = []string{
	"R1R29750",
	"R1R29751",
	"R0S29752",
	"R1R29753",
	"R0R29754",
	"R1R30027",
	"R1R30028",
}
var functionalProjects = []string{
	"RDX0000A",
	"RDX0000H",
	"RDX0000S",
	"RDX0000T",
	"RDX000PT",
	"X0000T",
}

func parseProjectID(projectName string) string {
	r, _ := regexp.Compile(".*\\((.*)\\)$")
	if r != nil {
		g := r.FindStringSubmatch(projectName)
		if len(g) > 1 {
			return g[1]
		}
	}
	return "N/A"
}
