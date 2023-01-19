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
		e, ok := elements[0].(string)
		if !ok {
			return "", pivot.InvalidType(elements[0])
		}
		var prefix string
		if prefixProject {
			prefix = e + "-"
		}
		teamWorkload, ok := projectsTeamWorkload[e]
		if ok {
			for _, p := range teamWorkload.budget {
				if p == elements[1] && Workload == elements[2] {
					return prefix + "Budget", nil
				}
			}
			for _, p := range teamWorkload.extension {
				if p == elements[1] && Workload == elements[2] {
					return prefix + "Extension", nil
				}
			}
			for _, p := range teamWorkload.other {
				if p == elements[1] && Workload == elements[2] {
					return prefix + "Other", nil
				}
			}
		}
		otherCosts, ok := projectsOtherCosts[e]
		if ok {
			for _, p := range otherCosts.budget {
				if p == elements[2] {
					return prefix + "Budget", nil
				}
			}
			for _, p := range otherCosts.extension {
				if p == elements[2] {
					return prefix + "Extension", nil
				}
			}
			for _, p := range otherCosts.other {
				if p == elements[2] {
					return prefix + "Other", nil
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
							fmt.Printf("WARNING: duplicated people detected: %q\n", p)
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
// record[6]=hours
// record[9]=project
// record[10]=task
// record[12-17]=hours(weekly)
// -->
// record[0]=project
// record[1]=task
// record[2]=employee
// record[3]=manager
// record[4-15]=hours(monthly)
func groupEBSTimeCardsByMonth(csvData [][]string, verbose bool) ([][]interface{}, error) {
	tams := map[string]*timeAndMaterial{}
	for _, record := range csvData {
		project := parseProjectID(record[9])
		employee := strings.TrimSpace(record[1])
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
		tam.AddWorkload(record[10], employee, record[0], month, monthHours, nextMonthHours, 0)
	}

	var fillRecord = func(employee, manager, project, task string, hours []float64) []interface{} {
		record := make([]interface{}, 16)
		record[0] = project
		record[1] = task
		record[2] = employee
		record[3] = manager
		record[4] = hours[0]
		record[5] = hours[1]
		record[6] = hours[2]
		record[7] = hours[3]
		record[8] = hours[4]
		record[9] = hours[5]
		record[10] = hours[6]
		record[11] = hours[7]
		record[12] = hours[8]
		record[13] = hours[9]
		record[14] = hours[10]
		record[15] = hours[11]
		return record
	}

	var rawData [][]interface{}
	for project, tam := range tams {
		for task, entry := range tam.entries {
			for employee, workload := range entry.workload {
				rawData = append(rawData, fillRecord(employee, workload.manager, project, task, workload.hours))
			}
		}
	}
	return rawData, nil
}
