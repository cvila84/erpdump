package ebs

import (
	"github.com/cvila84/erpdump/pkg/pivot"
)

func GenerateFromEBSExport(csvData string, csvTable string) error {
	rawData, err := readCsvFile(csvData)
	if err != nil {
		return err
	}
	pivotData, err := groupEBSTimeCardsByMonth(rawData)
	if err != nil {
		return err
	}
	// record[0]=employee
	// record[1]=manager
	// record[2]=project
	// record[3]=task
	// record[4-15]=hours(monthly)
	table := pivot.NewFloatTable(pivotData).
		//Filter(1, table.In(otaManagers)).
		Filter(2, pivot.In(otaProjects)).
		//Row([]int{0}, table.Group([][]string{OtaPeople}, []string{"OTA"}, "External"), nil, table.AlphaSort).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		Column([]int{2}, pivot.Group([][]string{otaProjects, functionalProjects}, []string{"OTA", "Functional"}, "Other"), nil, pivot.AlphaSort).
		StandardColumn(2).
		//StandardColumn(3).
		Values([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, pivot.YearlyHours, pivot.Sum, nil)
	err = table.Generate()
	if err != nil {
		return err
	}
	return saveCsvFile(csvTable, table.ToCSV())
}

func GenerateFromFinanceExport(csvData string, csvTable string) error {
	rawData, err := readCsvFile(csvData)
	if err != nil {
		return err
	}
	pivotData, err := filterBudgetPivotData(rawData)
	if err != nil {
		return err
	}
	// record[0]=employee
	// record[1]=project
	// record[2]=category
	// record[3-14]=hours
	// record[15-26]=cost
	table := pivot.NewFloatTable(pivotData).
		Row([]int{1}, nil, nil, pivot.AlphaSort).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		StandardColumn(2).
		Values([]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, pivot.YearlyHours, pivot.Sum, nil)
	err = table.Generate()
	if err != nil {
		return err
	}
	return saveCsvFile(csvTable, table.ToCSV())
}
