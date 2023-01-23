package ebs

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
	"github.com/cvila84/erpdump/pkg/utils"
	"golang.org/x/text/encoding/charmap"
	"os"
	"regexp"
	"strings"
)

var monthlySplit pivot.Compute[string] = func(elements []pivot.RawValue) (string, error) {
	e, ok := elements[0].(string)
	if !ok {
		return "", pivot.InvalidType(elements[0])
	}
	month, _, err := utils.ParseDateYYYYsMM(e)
	if err != nil {
		return "", fmt.Errorf("invalid YYYY-MM format for element %q: %w", e, err)
	}
	return utils.Month(month), nil
}

var quaterlySplit pivot.Compute[string] = func(elements []pivot.RawValue) (string, error) {
	e, ok := elements[0].(string)
	if !ok {
		return "", pivot.InvalidType(elements[0])
	}
	month, _, err := utils.ParseDateYYYYsMM(e)
	if err != nil {
		return "", fmt.Errorf("invalid YYYY-MM format for element %q: %w", e, err)
	}
	return utils.Quarter(month), nil
}

var dailyRate pivot.Compute[float64] = func(elements []pivot.RawValue) (float64, error) {
	hours, ok := elements[0].(float64)
	if !ok {
		return 0, pivot.InvalidType(elements[0])
	}
	cost, ok := elements[1].(float64)
	if !ok {
		return 0, pivot.InvalidType(elements[0])
	}
	if hours == 0 {
		return 0, nil
	} else {
		return -8 * cost / hours, nil
	}
}

var projectGroups = func(prefixProject bool) pivot.Compute[string] {
	return func(elements []pivot.RawValue) (string, error) {
		project, ok := elements[0].(string)
		if !ok {
			return "", pivot.InvalidType(elements[0])
		}
		var prefix string
		if prefixProject {
			prefix = project + "-"
		}
		teamWorkload, ok := projectsWorkloadSplit[project]
		if ok {
			for k, v := range teamWorkload {
				for _, p := range v {
					if p == elements[1] && Workload == elements[2] {
						return prefix + k, nil
					}
				}
			}
		}
		otherCosts, ok := projectOtherCostsSplit[project]
		if ok {
			for k, v := range otherCosts {
				for _, p := range v {
					if p == elements[2] {
						return prefix + k, nil
					}
				}
			}
		}
		return prefix + "Unknown", nil
	}
}

func uniquePeople(verbose bool, index int, peopleLists ...[][]string) []string {
	var result []string
	for _, l1 := range peopleLists {
		for _, l2 := range l1 {
			if len(l2[index]) > 0 {
				present := false
				for _, p := range result {
					if l2[index] == p {
						present = true
						if verbose {
							fmt.Printf("WARNING: duplicated people %q detected in %v\n", p, l1)
						}
					}
				}
				if !present {
					result = append(result, l2[index])
				}
			}
		}
	}
	return result
}

func parseProjectID(projectName string) string {
	r, _ := regexp.Compile(".*\\((.*)\\)$")
	if r != nil {
		g := r.FindStringSubmatch(projectName)
		if len(g) > 1 {
			return g[1]
		}
	}
	return "N/A"
}

func filesToRawData(csvDataFiles []string) ([][]interface{}, error) {
	var rawData [][]interface{}
	for i, csvDataFile := range csvDataFiles {
		data, err := readCsvFile(csvDataFile)
		if err != nil {
			return nil, fmt.Errorf("while reading %q: %w", csvDataFile, err)
		}
		for j, record := range data {
			if i == 0 || j > 0 {
				rawRecord := make([]interface{}, len(record))
				for j := 0; j < len(record); j++ {
					rawRecord[j] = record[j]
				}
				rawData = append(rawData, rawRecord)
			}
		}
	}
	return rawData, nil
}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(charmap.Windows1252.NewDecoder().Reader(f))
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func saveCsvFile(filePath string, csvData string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(csvData)
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

// groupEBSTimeCardsByMonth
// record[0]=manager
// record[1]=employee
// record[6]=date
// record[9]=project
// record[10]=task
// record[12-17]=hours(weekly)
// -->
// record[0]=project
// record[1]=task
// record[2]=employee
// record[3]=manager
// record[4]=month (yyyy-mm)
// record[5]=hours
func groupEBSTimeCardsByMonth(csvData [][]interface{}, verbose bool) ([][]interface{}, error) {
	tams := map[string]*timeAndMaterial{}
	for _, record := range csvData {
		project := parseProjectID(record[9].(string))
		employee := strings.TrimSpace(record[1].(string))
		month, monthHours, nextMonthHours, err := monthlyHours(record)
		if err != nil {
			return nil, fmt.Errorf("cannot parse week hour fields %v: %w", record, err)
		}
		if monthHours == 0 && nextMonthHours == 0 {
			if verbose {
				fmt.Printf("WARNING: no computed hours for entry %v\n", record)
			}
			continue
		}
		tam, ok := tams[project]
		if !ok {
			tam = &timeAndMaterial{}
			tams[project] = tam
		}
		tam.AddWorkload(record[10].(string), employee, record[0].(string), month, monthHours, nextMonthHours, 0)
	}

	var rawData [][]interface{}
	for project, tam := range tams {
		for task, entry := range tam.entries {
			for employee, workload := range entry.workload {
				for month, hours := range workload.hours {
					rawData = append(rawData, []interface{}{project, task, employee, workload.manager, utils.ToDateYYYYsMM(month+1, 2022), hours})
				}
			}
		}
	}
	return rawData, nil
}
