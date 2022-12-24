package erp

type MonthlyHours []float64

func newMonthlyHours() MonthlyHours {
	return make(MonthlyHours, 12)
}

type ProjectTime map[string]MonthlyHours

func newProjectTime() ProjectTime {
	return make(ProjectTime)
}

type EmployeeTimes struct {
	Name  Person
	times map[string]ProjectTime
}

func (e *EmployeeTimes) Add(projectName string, taskName string, month int, hoursInMonth float64, hoursInNextMonth float64) {
	if e.times == nil {
		e.times = make(map[string]ProjectTime)
	}
	project, ok := e.times[projectName]
	if !ok {
		project = newProjectTime()
		e.times[projectName] = project
	}
	taskTime, ok := project[taskName]
	if !ok {
		taskTime = newMonthlyHours()
		project[taskName] = taskTime
	}
	taskTime[month-1] += hoursInMonth
	if hoursInNextMonth > 0 {
		taskTime[month] += hoursInNextMonth
	}
}
