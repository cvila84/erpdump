package ebs

import (
	"github.com/cvila84/erpdump/pkg/utils"
	"strconv"
)

type projectTAM struct {
	category string
	hours    map[string][]float64
	costs    []float64
}

type timeAndMaterial struct {
	employee string
	manager  string
	projects map[string]*projectTAM
}

func (t *timeAndMaterial) AddHours(projectName, taskName, category string, month int, hoursInMonth, hoursInNextMonth float64) {
	if t.projects == nil {
		t.projects = make(map[string]*projectTAM)
	}
	project, ok := t.projects[projectName]
	if !ok {
		project = &projectTAM{
			category: category,
			hours:    make(map[string][]float64),
			costs:    make([]float64, 12),
		}
		t.projects[projectName] = project
	}
	taskTime, ok := project.hours[taskName]
	if !ok {
		taskTime = make([]float64, 12)
		project.hours[taskName] = taskTime
	}
	taskTime[month-1] += hoursInMonth
	if hoursInNextMonth > 0 {
		taskTime[month] += hoursInNextMonth
	}
}

func (t *timeAndMaterial) AddCosts(projectName, category string, month int, costsInMonth float64) {
	if t.projects == nil {
		t.projects = make(map[string]*projectTAM)
	}
	project, ok := t.projects[projectName]
	if !ok {
		project = &projectTAM{
			category: category,
			hours:    make(map[string][]float64),
			costs:    make([]float64, 12),
		}
		t.projects[projectName] = project
	}
	project.costs[month-1] += costsInMonth
}

func weeklyHours(record []string) ([]float64, error) {
	var hours []float64
	for j := 12; j < 17; j++ {
		var hour float64
		var err error
		if len(record[j]) > 0 {
			hour, err = strconv.ParseFloat(record[j], 32)
			if err != nil {
				return nil, err
			}
		} else {
			hour = 0
		}
		hours = append(hours, hour)
	}
	return hours, nil
}

func monthlyHours(record []string) (int, float64, float64, error) {
	startDay, startMonth, startYear, err := utils.ParseDateDDsMMMsYYYY(record[6])
	if err != nil {
		return 0, 0, 0, err
	}
	daysInMonth := utils.DaysIn(startMonth, startYear)
	var hoursInMonth float64
	var hoursInNextMonth float64
	hours, err := weeklyHours(record)
	if err != nil {
		return 0, 0, 0, err
	}
	for i := 0; i < 5; i++ {
		if startDay+i <= daysInMonth {
			hoursInMonth += hours[i]
		} else {
			hoursInNextMonth += hours[i]
		}
	}
	return startMonth, hoursInMonth, hoursInNextMonth, nil
}
