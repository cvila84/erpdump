package utils

import (
	"strconv"
	"time"
)

func Quarter(month int) string {
	switch month {
	case 1:
		return "Q1"
	case 2:
		return "Q1"
	case 3:
		return "Q1"
	case 4:
		return "Q2"
	case 5:
		return "Q2"
	case 6:
		return "Q2"
	case 7:
		return "Q3"
	case 8:
		return "Q3"
	case 9:
		return "Q3"
	case 10:
		return "Q4"
	case 11:
		return "Q4"
	case 12:
		return "Q4"
	}
	return "Q?" + strconv.Itoa(month)
}

func Month(month int) string {
	switch month {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	case 4:
		return "Apr"
	case 5:
		return "May"
	case 6:
		return "Jun"
	case 7:
		return "Jul"
	case 8:
		return "Aug"
	case 9:
		return "Sep"
	case 10:
		return "Oct"
	case 11:
		return "Nov"
	case 12:
		return "Dec"
	}
	return "M?" + strconv.Itoa(month)
}

func ParseDateYYYYsMM(date string) (int, int, error) {
	month, err := strconv.Atoi(date[5:7])
	if err != nil {
		return 0, 0, err
	}
	year, err := strconv.Atoi(date[0:4])
	if err != nil {
		return 0, 0, err
	}
	return month, year, nil
}

func ParseDateDDsMMMsYYYY(date string) (int, int, int, error) {
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
