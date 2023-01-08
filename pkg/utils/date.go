package utils

import (
	"strconv"
	"time"
)

func ParseDate(date string) (int, int, int, error) {
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

func DaysIn(m int, year int) int {
	return time.Date(year, time.Month(m+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
