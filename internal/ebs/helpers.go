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
	"strconv"
	"strings"
)

var float pivot.Compute[float64] = func(elements []interface{}) (float64, error) {
	return toFloat(elements[0])
}

var monthlySplit pivot.Compute[string] = func(elements []interface{}) (string, error) {
	e, ok := elements[0].(string)
	if !ok {
		return "", fmt.Errorf("invalid type %T for element %s", elements[0], elements[0])
	}
	month, _, err := utils.ParseDateYYYYsMM(e)
	if err != nil {
		return "", fmt.Errorf("invalid YYYY-MM format for element %s", e)
	}
	return utils.Month(month), nil
}

var quaterlySplit pivot.Compute[string] = func(elements []interface{}) (string, error) {
	e, ok := elements[0].(string)
	if !ok {
		return "", fmt.Errorf("invalid type %T for element %s", elements[0], elements[0])
	}
	month, _, err := utils.ParseDateYYYYsMM(e)
	if err != nil {
		return "", fmt.Errorf("invalid YYYY-MM format for element %s", e)
	}
	return utils.Quarter(month), nil
}

var dailyRate pivot.Compute[float64] = func(elements []interface{}) (float64, error) {
	hours, err := toFloat(elements[0])
	if err != nil {
		return 0, err
	}
	cost, err := toFloat(elements[1])
	if err != nil {
		return 0, err
	}
	if hours == 0 {
		return 0, nil
	} else {
		return -8 * cost / hours, nil
	}
}

var projectGroups = func(prefixProject bool) pivot.Compute[string] {
	return func(elements []interface{}) (string, error) {
		e, ok := elements[0].(string)
		if !ok {
			return "", fmt.Errorf("invalid type %T for element %s", elements[0], elements[0])
		}
		var prefix string
		if prefixProject {
			prefix = e + "-"
		}
		team, ok := projectsTeam[e]
		if ok {
			for _, p := range team.budget {
				if p == elements[1] {
					return prefix + "Budget", nil
				}
			}
			for _, p := range team.extension {
				if p == elements[1] {
					return prefix + "Ext", nil
				}
			}
			for _, p := range team.other {
				if p == elements[1] {
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

func toFloat(element interface{}) (float64, error) {
	e, ok := element.(string)
	if !ok {
		return 0, fmt.Errorf("invalid type %T for element %s", element, element)
	}
	e = strings.Replace(e, ",", ".", 1)
	result, err := strconv.ParseFloat(e, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid numeric format for element %s", e)
	}
	return result, nil
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

// filterBudgetPivotData
// record[3]=month (yyyy-mm)
// record[14]=project
// record[21]=cost
// record[26]=category
// record[32]=employee
// record[40]=hours
// -->
// record[0]=employee
// record[1]=project
// record[2]=category
// record[3-14]=hours
// record[15-26]=cost

func filterBudgetPivotData(csvData [][]string, verbose bool) ([][]interface{}, error) {
	tams := map[string]*timeAndMaterial{}
	for _, record := range csvData {
		project := strings.TrimSpace(record[14])
		employee := strings.TrimSpace(record[32])
		month, _, err := utils.ParseDateYYYYsMM(record[3])
		if err != nil {
			return nil, fmt.Errorf("cannot parse month field %q: %w", record[3], err)
		}
		record[40] = strings.Replace(record[40], ",", ".", 1)
		monthHours, err := strconv.ParseFloat(record[40], 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse hours field %q: %w", record[40], err)
		}
		record[21] = strings.Replace(record[21], ",", ".", 1)
		monthCost, err := strconv.ParseFloat(record[21], 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse hours field %q: %w", record[21], err)
		}
		monthCost = -monthCost
		if monthHours == 0 && monthCost == 0 {
			if verbose {
				fmt.Printf("WARNING: no computed hours nor costs for entry %v\n", record)
			}
			continue
		}
		tam, ok := tams[project]
		if !ok {
			tam = &timeAndMaterial{}
			tams[project] = tam
		}
		category := strings.TrimSpace(record[26])
		if strings.ToUpper(category) == "WORKLOAD" {
			tam.AddWorkload("", employee, "", month, monthHours, 0, monthCost)
		} else {
			tam.AddCosts("", category, month, monthCost)
		}
	}

	var fillRecord = func(employee, project, category string, hours, costs []float64) []interface{} {
		record := make([]interface{}, 27)
		if hours == nil {
			hours = make([]float64, 12)
		}
		record[0] = employee
		record[1] = project
		record[2] = category
		record[3] = hours[0]
		record[4] = hours[1]
		record[5] = hours[2]
		record[6] = hours[3]
		record[7] = hours[4]
		record[8] = hours[5]
		record[9] = hours[6]
		record[10] = hours[7]
		record[11] = hours[8]
		record[12] = hours[9]
		record[13] = hours[10]
		record[14] = hours[11]
		record[15] = costs[0]
		record[16] = costs[1]
		record[17] = costs[2]
		record[18] = costs[3]
		record[19] = costs[4]
		record[20] = costs[5]
		record[21] = costs[6]
		record[22] = costs[7]
		record[23] = costs[8]
		record[24] = costs[9]
		record[25] = costs[10]
		record[26] = costs[11]
		return record
	}

	var rawData [][]interface{}
	for project, tam := range tams {
		entry := tam.entries[""]
		for employee, workload := range entry.workload {
			rawData = append(rawData, fillRecord(employee, project, "Workload", workload.hours, workload.costs))
		}
		for category, costs := range entry.others {
			rawData = append(rawData, fillRecord("N/A", project, category, nil, costs))
		}
	}

	var sb strings.Builder
	for _, r := range rawData {
		sb.WriteString(fmt.Sprintf("%q;%q;%q;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f;%.2f\n", r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8], r[9], r[10], r[11], r[12], r[13], r[14], r[15], r[16], r[17], r[18], r[19], r[20], r[21], r[22], r[23], r[24], r[25], r[26]))
	}
	_ = saveCsvFile("./filterBudgetPivotData-debug.csv", sb.String())
	return rawData, nil
}
