package ebs

import (
	"github.com/cvila84/erpdump/pkg/utils"
	"strconv"
)

type tamWorkload struct {
	manager string
	hours   []float64
	costs   []float64
}

type tamEntry struct {
	workload map[string]*tamWorkload
	others   map[string][]float64
}

type timeAndMaterial struct {
	entries map[string]*tamEntry
}

func (t *timeAndMaterial) AddWorkload(taskName, employee, manager string, month int, hoursInMonth, hoursInNextMonth, costsInMonth float64) {
	if t.entries == nil {
		t.entries = make(map[string]*tamEntry)
	}
	entry, ok := t.entries[taskName]
	if !ok {
		entry = &tamEntry{
			workload: make(map[string]*tamWorkload),
		}
		t.entries[taskName] = entry
	}
	workload, ok := entry.workload[employee]
	if !ok {
		workload = &tamWorkload{
			manager: manager,
			hours:   make([]float64, 12),
			costs:   make([]float64, 12),
		}
		entry.workload[employee] = workload
	}
	workload.hours[month-1] += hoursInMonth
	if hoursInNextMonth > 0 {
		workload.hours[month] += hoursInNextMonth
	}
	workload.costs[month-1] += costsInMonth
}

func (t *timeAndMaterial) AddCosts(taskName, category string, month int, costsInMonth float64) {
	if t.entries == nil {
		t.entries = make(map[string]*tamEntry)
	}
	entry, ok := t.entries[taskName]
	if !ok {
		entry = &tamEntry{
			workload: make(map[string]*tamWorkload),
		}
		t.entries[taskName] = entry
	}
	if entry.others == nil {
		entry.others = make(map[string][]float64)
	}
	costs, ok := entry.others[category]
	if !ok {
		costs = make([]float64, 12)
		entry.others[category] = costs
	}
	costs[month-1] += costsInMonth
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
