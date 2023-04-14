package ebs

import (
	"fmt"
	"github.com/cvila84/pivot"
	"path/filepath"
)

func GenerateFromEBSExport(csvDataFiles []string, csvTablePath, csvTablePrefix string, verbose bool) error {
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
	err = table.Generate(verbose)
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
	err = table.Generate(verbose)
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
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-ctr.csv", table.ToCSV())
	if err != nil {
		return err
	}

	return nil
}

type FinanceRecordKind int

const (
	Month FinanceRecordKind = iota
	Project
	Cost
	Category
	ExpOrg
	Employee
	Hours
)

type FinanceRecordIndexes map[FinanceRecordKind]int

var RecordIndexes2022 = FinanceRecordIndexes{
	Month:    3,
	Project:  14,
	Cost:     21,
	Category: 26,
	ExpOrg:   27,
	Employee: 32,
	Hours:    40,
}

var RecordIndexes2023 = FinanceRecordIndexes{
	Month:    6,
	Project:  17,
	Cost:     24,
	Category: 29,
	ExpOrg:   30,
	Employee: 36,
	Hours:    44,
}

func GenerateFromFinanceExport(csvDataFiles []string, csvRecordIndexes FinanceRecordIndexes, csvTablePath, csvTablePrefix string, verbose bool) error {
	pivotData, err := filesToRawData([]string{csvDataFiles[0]})
	if err != nil {
		return fmt.Errorf("while reading %v: %w", csvDataFiles, err)
	}

	table := pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Hours], pivot.Sum, pivot.Digits(0)).
		Values(csvRecordIndexes[Cost], pivot.Sum, pivot.Digits(0)).
		ComputedValues("DailyRate", pivot.DataRefs([]int{csvRecordIndexes[Hours], csvRecordIndexes[Cost]}, pivot.Sum), dailyRate, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rd.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Hours], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdh1.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Hours], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdh2.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Cost], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdc1.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Cost], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-rdc2.csv", table.ToCSV())
	if err != nil {
		return err
	}

	pivotData, err = filesToRawData([]string{csvDataFiles[1]})
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Hours], pivot.Sum, pivot.Digits(0)).
		Values(csvRecordIndexes[Cost], pivot.Sum, pivot.Digits(0)).
		ComputedValues("DailyRate", pivot.DataRefs([]int{csvRecordIndexes[Hours], csvRecordIndexes[Cost]}, pivot.Sum), dailyRate, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Hours], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3h.csv", table.ToCSV())
	if err != nil {
		return err
	}

	table = pivot.NewTable(pivotData, true).
		ComputedRow([]int{csvRecordIndexes[Project]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Project], csvRecordIndexes[Employee], csvRecordIndexes[Category]}, nil, projectPeopleGroups(false), pivot.AlphaSort).
		Row(csvRecordIndexes[Category]).
		ComputedRow([]int{csvRecordIndexes[ExpOrg]}, nil, nil, pivot.AlphaSort).
		ComputedRow([]int{csvRecordIndexes[Employee]}, nil, nil, pivot.AlphaSort).
		ComputedColumn([]int{csvRecordIndexes[Month]}, nil, monthlySplit, pivot.MonthSort).
		Values(csvRecordIndexes[Cost], pivot.Sum, pivot.Digits(0))
	err = table.Generate(verbose)
	if err != nil {
		return err
	}
	err = saveCsvFile(csvTablePath+string(filepath.Separator)+csvTablePrefix+"-l3c.csv", table.ToCSV())
	if err != nil {
		return err
	}

	return nil
}
