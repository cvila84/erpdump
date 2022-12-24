package utils

import (
	"strconv"
	"time"
)

func parseDate(date string) (int, int, int, error) {
	day, err := strconv.Atoi(date[0:2])
	if err != nil {
		return 0, 0, 0, err
	}
	sMonth := date[3:6]
	month := 0
	switch {
	case sMonth == "JAN":
		month = 1
	case sMonth == "FEB":
		month = 2
	case sMonth == "MAR":
		month = 3
	case sMonth == "APR":
		month = 4
	case sMonth == "MAY":
		month = 5
	case sMonth == "JUN":
		month = 6
	case sMonth == "JUL":
		month = 7
	case sMonth == "AUG":
		month = 8
	case sMonth == "SEP":
		month = 9
	case sMonth == "OCT":
		month = 10
	case sMonth == "NOV":
		month = 11
	case sMonth == "DEC":
		month = 12
	}
	year, err := strconv.Atoi(date[7:11])
	if err != nil {
		return 0, 0, 0, err
	}
	return day, month, year, nil
}

func daysIn(m int, year int) int {
	return time.Date(year, time.Month(m+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

func WeeklyHours(record []string) []float64 {
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

func MonthlyHours(record []string) (int, float64, float64, error) {
	startDay, startMonth, startYear, err := parseDate(record[6])
	if err != nil {
		return 0, 0, 0, err
	}
	daysInMonth := daysIn(startMonth, startYear)
	var hoursInMonth float64
	var hoursInNextMonth float64
	weeklyHours := WeeklyHours(record)
	for i := 0; i < 5; i++ {
		if startDay+i <= daysInMonth {
			hoursInMonth += weeklyHours[i]
		} else {
			hoursInNextMonth += weeklyHours[i]
		}
	}
	return startMonth, hoursInMonth, hoursInNextMonth, nil
}
