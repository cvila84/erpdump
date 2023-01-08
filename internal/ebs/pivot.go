package ebs

import (
	"bufio"
	"github.com/cvila84/erpdump/pkg/pivot"
	"os"
)

func GenerateFromEBSExport(csvData string, csvTable string) error {
	weeklyTimeCards, err := readCsvFile(csvData)
	if err != nil {
		return err
	}
	monthlyData, err := GroupEBSTimeCardsByMonth(weeklyTimeCards)
	if err != nil {
		return err
	}
	// record[0]=employee
	// record[1]=manager
	// record[2]=project
	// record[3]=task
	// record[4-15]=hours(monthly)
	table := pivot.NewFloatTable(monthlyData).
		//Filter(1, table.In(otaManagers)).
		Filter(2, pivot.In(OtaProjects)).
		//Row([]int{0}, table.Group([][]string{OtaPeople}, []string{"OTA"}, "External"), nil, table.AlphaSort).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		Column([]int{2}, pivot.Group([][]string{OtaProjects, FunctionalProjects}, []string{"OTA", "Functional"}, "Other"), nil, pivot.AlphaSort).
		StandardColumn(2).
		//StandardColumn(3).
		Values([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, pivot.YearlyHours, pivot.Sum, nil)
	err = table.Generate()
	if err != nil {
		return err
	}
	file, err := os.Create(csvTable)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(table.ToCSV())
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func GenerateFromFinanceExport(csvData string, csvTable string) error {
	//budgetPivotRawData, err := readCsvFile(csvData)
	//if err != nil {
	//	return err
	//}
	//monthlyData, err := FilterBudgetPivotData(budgetPivotRawData)
	//if err != nil {
	//	return err
	//}
	return nil
}
