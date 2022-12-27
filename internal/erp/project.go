package erp

import (
	"github.com/cvila84/erpdump/pkg/utils"
	"regexp"
)

type Project struct {
	Name  string
	ID    string
	tasks *utils.Vector[string]
}

func (p *Project) AddTask(taskName string) {
	if p.tasks == nil {
		p.tasks = &utils.Vector[string]{ID: func(element string) string { return element }}
	}
	p.tasks.Add(taskName)
}

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

func NewProject(projectName string) Project {
	return Project{Name: projectName, ID: ParseProjectID(projectName)}
}
