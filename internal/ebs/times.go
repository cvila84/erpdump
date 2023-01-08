package ebs

import (
	"github.com/cvila84/erpdump/pkg/utils"
	"strconv"
)

type TaskMonthlyHours []float64

func newTaskMonthlyHours() TaskMonthlyHours {
	return make(TaskMonthlyHours, 12)
}

type ProjectTimes map[string]TaskMonthlyHours

func newProjectTimes() ProjectTimes {
	return make(ProjectTimes)
}

type EmployeeTimes struct {
	Name        string
	ManagerName string
	times       map[string]ProjectTimes
}

func (e *EmployeeTimes) Add(projectName string, taskName string, month int, hoursInMonth float64, hoursInNextMonth float64) {
	if e.times == nil {
		e.times = make(map[string]ProjectTimes)
	}
	project, ok := e.times[projectName]
	if !ok {
		project = newProjectTimes()
		e.times[projectName] = project
	}
	taskTime, ok := project[taskName]
	if !ok {
		taskTime = newTaskMonthlyHours()
		project[taskName] = taskTime
	}
	taskTime[month-1] += hoursInMonth
	if hoursInNextMonth > 0 {
		taskTime[month] += hoursInNextMonth
	}
}

func (e *EmployeeTimes) GetAll() [][]interface{} {
	var result [][]interface{}
	for k1, v1 := range e.times {
		for k2, v2 := range v1 {
			record := make([]interface{}, 16)
			result = append(result, record)
			record[0] = e.Name
			record[1] = e.ManagerName
			record[2] = k1
			record[3] = k2
			record[4] = v2[0]
			record[5] = v2[1]
			record[6] = v2[2]
			record[7] = v2[3]
			record[8] = v2[4]
			record[9] = v2[5]
			record[10] = v2[6]
			record[11] = v2[7]
			record[12] = v2[8]
			record[13] = v2[9]
			record[14] = v2[10]
			record[15] = v2[11]
		}
	}
	return result
}

func weeklyHours(record []string) []float64 {
	var hours []float64
	for j := 12; j < 17; j++ {
		var hour float64
		var err error
		if len(record[j]) > 0 {
			hour, err = strconv.ParseFloat(record[j], 32)
			if err != nil {
				panic(err)
			}
		} else {
			hour = 0
		}
		hours = append(hours, hour)
	}
	return hours
}

func monthlyHours(record []string) (int, float64, float64, error) {
	startDay, startMonth, startYear, err := utils.ParseDate(record[6])
	if err != nil {
		return 0, 0, 0, err
	}
	daysInMonth := utils.DaysIn(startMonth, startYear)
	var hoursInMonth float64
	var hoursInNextMonth float64
	hours := weeklyHours(record)
	for i := 0; i < 5; i++ {
		if startDay+i <= daysInMonth {
			hoursInMonth += hours[i]
		} else {
			hoursInNextMonth += hours[i]
		}
	}
	return startMonth, hoursInMonth, hoursInNextMonth, nil
}
