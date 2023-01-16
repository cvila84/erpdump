package ebs

import (
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
	"path/filepath"
)

func GenerateFromEBSExport(csvDataFile, csvTablePath, csvTablePrefix string) error {
	rawData, err := readCsvFile(csvDataFile)
	if err != nil {
		return fmt.Errorf("while reading %q: %w", csvDataFile, err)
	}
	// record[0]=manager
	// record[1]=employee
	// record[6]=hours
	// record[9]=project
	// record[10]=task
	// record[12-17]=hours(weekly)

	pivotData, err := groupEBSTimeCardsByMonth(rawData[1:], false)
	if err != nil {
		return fmt.Errorf("while processing raw csv data: %w", err)
	}
	// record[0]=project
	// record[1]=task
	// record[2]=employee
	// record[3]=manager
	// record[4-15]=hours(monthly)

	table := pivot.NewTable(pivotData, false).
		//Filter(1, table.In(otaManagers)).
		Filter(2, pivot.In(otaProjects)).
		//Row([]int{0}, table.Group([][]string{OtaPeople}, []string{"OTA"}, "External"), nil, table.AlphaSort).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		Column([]int{2}, nil, pivot.Group([][]string{otaProjects, functionalProjects}, []string{"OTA", "Functional"}, "Other"), pivot.AlphaSort).
		StandardColumn(2).
		//StandardColumn(3).
		Values([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, pivot.SumFloats, pivot.Sum, pivot.Digits(1))
	err = table.Generate()
	if err != nil {
		return err
	}
	return saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+".csv", table.ToCSV())
}

func GenerateFromFinanceExport(csvDataFiles []string, csvTablePath, csvTablePrefix string) error {
	var allRawData [][]interface{}
	for i, csvDataFile := range csvDataFiles {
		rawData, err := readCsvFile(csvDataFile)
		if err != nil {
			return fmt.Errorf("while reading %q: %w", csvDataFile, err)
		}
		for j, record := range rawData {
			if i == 0 || j > 0 {
				rawRecord := make([]interface{}, len(record))
				for j := 0; j < len(record); j++ {
					rawRecord[j] = record[j]
				}
				allRawData = append(allRawData, rawRecord)
			}
		}
	}
	// record[3]=month (yyyy-mm)
	// record[14]=project
	// record[21]=cost
	// record[26]=category
	// record[32]=employee
	// record[40]=hours

	table := pivot.NewTable(allRawData, true).
		Row([]int{14}, nil, nil, pivot.AlphaSort).
		Row([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		StandardRow(26).
		Row([]int{32}, nil, nil, pivot.AlphaSort).
		Column([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values([]int{40, 21}, dailyRate, pivot.Set, pivot.Digits(1)).
		Values([]int{40}, nil, pivot.Sum, pivot.Digits(0)).
		Values([]int{21}, nil, pivot.Sum, pivot.Digits(0))
	err := table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+".csv", table.ToCSV())
	if err != nil {
		return err
	}
	return nil
}
