package ebs

import (
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
	"path/filepath"
)

func GenerateFromEBSExport(csvDataFiles []string, csvTablePath, csvTablePrefix string) error {
	data, err := filesToRawData(csvDataFiles)
	if err != nil {
		return fmt.Errorf("while reading %v: %w", csvDataFiles, err)
	}
	// record[0]=manager
	// record[1]=employee
	// record[6]=date
	// record[9]=project
	// record[10]=task
	// record[12-17]=hours(weekly)

	pivotData, err := groupEBSTimeCardsByMonth(data[1:], false)
	if err != nil {
		return fmt.Errorf("while processing raw csv data: %w", err)
	}
	// record[0]=project
	// record[1]=task
	// record[2]=employee
	// record[3]=manager
	// record[4]=month (yyyy-mm)
	// record[5]=hours

	projectDevGroups := pivot.Group(
		[][]string{otaDevProjects, otaCogsProjects, otaOcosProjects, tacDevProjects, iotDevProjects, otaOtherDevProjects, functionalOtherProjects, functionalHolidaysProjects},
		[]string{"COTA-Dev", "COTA-Custom", "COTA-L3", "TAC-Dev", "IOT-Dev", "OTA-Other", "Other", "Holidays"},
		"Non-OTA",
	)

	table := pivot.NewTable(pivotData, false).
		ComputedRow([]int{0}, nil, projectDevGroups, pivot.AlphaSort).
		ComputedRow([]int{0}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{3, 2}, nil, cotaManagerCountry, pivot.AlphaSort).
		ComputedRow([]int{2}, pivot.In(uniquePeople(false, 0, cotaDevL3BudgetPeople)), nil, pivot.AlphaSort).
		ComputedColumn([]int{4}, nil, monthlySplit, pivot.MonthSort).
		Values(5, pivot.Sum, pivot.Digits(1))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-prj.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, false).
		ComputedRow([]int{3, 2}, nil, cotaManagerCountry, pivot.AlphaSort).
		ComputedRow([]int{2}, pivot.In(uniquePeople(false, 0, cotaDevL3BudgetPeople)), nil, pivot.AlphaSort).
		ComputedRow([]int{0}, nil, projectDevGroups, pivot.AlphaSort).
		ComputedRow([]int{0}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{4}, nil, monthlySplit, pivot.MonthSort).
		Values(5, pivot.Sum, pivot.Digits(1))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-res.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, false).
		Filter(2, pivot.In(uniquePeople(false, 0, cotaDevL3BudgetPeople))).
		ComputedRow([]int{3, 2}, nil, cotaManagerCountry, pivot.AlphaSort).
		ComputedRow([]int{0}, nil, projectDevGroups, pivot.AlphaSort).
		ComputedRow([]int{0}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{4}, nil, monthlySplit, pivot.MonthSort).
		Values(5, pivot.Sum, pivot.Digits(1))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-ctr.csv", table.ToCSV())
	if err != nil {
		return err
	}

	return nil
}

func GenerateFromFinanceExport(csvDataFiles []string, csvTablePath, csvTablePrefix string) error {
	pivotData, err := filesToRawData([]string{csvDataFiles[0]})
	if err != nil {
		return fmt.Errorf("while reading %v: %w", csvDataFiles, err)
	}

	// record[3]=month (yyyy-mm)
	// record[14]=project
	// record[21]=cost
	// record[26]=category
	// record[27]=exporg
	// record[32]=employee
	// record[40]=hours

	table := pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(40, pivot.Sum, pivot.Digits(0)).
		Values(21, pivot.Sum, pivot.Digits(0)).
		ComputedValues("DailyRate", pivot.DataRefs([]int{40, 21}, pivot.Sum), dailyRate, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rd.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(40, pivot.Sum, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdh.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(21, pivot.Sum, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdc.csv", table.ToCSV())
	if err != nil {
		return err
	}

	pivotData, err = filesToRawData([]string{csvDataFiles[1]})
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(40, pivot.Sum, pivot.Digits(0)).
		Values(21, pivot.Sum, pivot.Digits(0)).
		ComputedValues("DailyRate", pivot.DataRefs([]int{40, 21}, pivot.Sum), dailyRate, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(40, pivot.Sum, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3h.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{14}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{14, 32, 26}, nil, projectGroups(false), pivot.AlphaSort).
		Row(26).
		ComputedRow([]int{27}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{32}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{3}, nil, monthlySplit, pivot.MonthSort).
		Values(21, pivot.Sum, pivot.Digits(0))
	err = table.Generate()
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3c.csv", table.ToCSV())
	if err != nil {
		return err
	}

	return nil
}
