package ebs

import (
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
	"path/filepath"
)

func GenerateFromEBSExport(csvDataFile, csvTablePath, csvTablePrefix string) error {
	rawData, err := readCsvFile(csvDataFile)
	if err != nil {
		return fmt.Errorf("while reading %s: %w", csvDataFile, err)
	}

	pivotData, err := groupEBSTimeCardsByMonth(rawData[1:])
	if err != nil {
		return fmt.Errorf("while processing raw csv data: %w", err)
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
	return saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+".csv", table.ToCSV())
}

func GenerateFromFinanceExport(csvDataFiles []string, csvTablePath, csvTablePrefix string) error {
	var allRawData [][]string
	for _, csvDataFile := range csvDataFiles {
		rawData, err := readCsvFile(csvDataFile)
		if err != nil {
			return fmt.Errorf("while reading %s: %w", csvDataFile, err)
		}
		allRawData = append(allRawData, rawData[1:]...)
	}

	pivotData, err := filterBudgetPivotData(allRawData)
	if err != nil {
		return fmt.Errorf("while processing raw csv data: %w", err)
	}
	// record[0]=employee
	// record[1]=project
	// record[2]=category
	// record[3-14]=hours
	// record[15-26]=cost

	//for _, r := range pivotData {
	//	fmt.Printf(
	//		"%q / %q / %q [ %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f ] [ %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f / %.2f ]\n",
	//		r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8], r[9], r[10], r[11], r[12], r[13], r[14], r[15], r[16], r[17], r[18], r[19], r[20], r[21], r[22], r[23], r[24], r[25], r[26],
	//	)
	//}

	table := pivot.NewFloatTable(pivotData).
		Row([]int{1}, nil, nil, pivot.AlphaSort).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		Column([]int{1, 0}, projectGroups(false), nil, pivot.AlphaSort).
		Values([]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, pivot.YearlyHours, pivot.Sum, nil)
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"h.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewFloatTable(pivotData).
		Row([]int{1}, nil, nil, pivot.AlphaSort).
		StandardRow(2).
		Row([]int{0}, nil, nil, pivot.AlphaSort).
		Column([]int{1, 0}, projectGroups(false), nil, pivot.AlphaSort).
		Values([]int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, pivot.YearlyHours, pivot.Sum, nil)
	err = table.Generate()
	if err != nil {
		return err
	}
	return saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"c.csv", table.ToCSV())
}
