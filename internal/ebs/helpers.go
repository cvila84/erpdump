package ebs

import (
	"bufio"
	"encoding/csv"
	"github.com/cvila84/erpdump/pkg/utils"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(f)
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
// record[0]=employee
// record[1]=manager
// record[2]=project
// record[3]=task
// record[4-15]=hours(monthly)
func groupEBSTimeCardsByMonth(csvData [][]string) ([][]interface{}, error) {
	tams := &utils.Vector[timeAndMaterial]{ID: func(element timeAndMaterial) string { return element.employee }}
	for i, record := range csvData {
		if i > 0 {
			employee := strings.TrimSpace(record[1])
			tam, ok := tams.Get(employee)
			if !ok {
				tam = &timeAndMaterial{employee: employee, manager: strings.TrimSpace(record[0])}
				tams.Add(tam)
			}
			month, monthHours, nextMonthHours, err := monthlyHours(record)
			if err != nil {
				return nil, err
			}
			if monthHours > 0 || nextMonthHours > 0 {
				tam.AddHours(parseProjectID(record[9]), record[10], month, monthHours, nextMonthHours)
			}
		}
	}
	var rawData [][]interface{}
	for _, tam := range tams.GetAll() {
		for k1, v1 := range tam.projects {
			for k2, v2 := range v1.hours {
				record := make([]interface{}, 16)
				rawData = append(rawData, record)
				record[0] = tam.employee
				record[1] = tam.manager
				record[2] = k1
				record[3] = k2
				record[4] = v2[0]
				record[5] = v2[1]
				record[6] = v2[2]
				record[7] = v2[3]
				record[8] = v2[4]
				record[9] = v2[5]
				record[10] = v2[6]
				record[11] = v2[7]
				record[12] = v2[8]
				record[13] = v2[9]
				record[14] = v2[10]
				record[15] = v2[11]
			}
		}
	}
	return rawData, nil
}

// filterBudgetPivotData
// record[3]=month (yyyy-mm)
// record[14]=project
// record[21]=cost
// record[28]=type
// record[32]=employee
// record[40]=hours
// -->
// record[0]=employee
// record[1]=project
// record[2-13]=hours
// record[14-25]=cost

func filterBudgetPivotData(csvData [][]string) ([][]interface{}, error) {
	tams := &utils.Vector[timeAndMaterial]{ID: func(element timeAndMaterial) string { return element.employee }}
	for i, record := range csvData {
		if i > 0 {
			employee := strings.TrimSpace(record[32])
			tam, ok := tams.Get(employee)
			if !ok {
				tam = &timeAndMaterial{employee: employee}
				tams.Add(tam)
			}
			month, _, err := utils.ParseDateYYYYsMM(record[3])
			if err != nil {
				return nil, err
			}
			record[40] = strings.Replace(record[40], ",", ".", 1)
			monthHours, err := strconv.ParseFloat(record[40], 32)
			if err != nil {
				return nil, err
			}
			record[21] = strings.Replace(record[21], ",", ".", 1)
			monthCost, err := strconv.ParseFloat(record[21], 32)
			if err != nil {
				return nil, err
			}
			if monthHours > 0 || monthCost > 0 {
				tam.AddHours(record[14], "", month, monthHours, 0)
				tam.AddCosts(record[14], month, monthCost)
			}
		}
	}
	var rawData [][]interface{}
	for _, tam := range tams.GetAll() {
		for k1, v1 := range tam.projects {
			record := make([]interface{}, 26)
			rawData = append(rawData, record)
			record[0] = tam.employee
			record[1] = k1
			record[2] = v1.hours[""][0]
			record[3] = v1.hours[""][1]
			record[4] = v1.hours[""][2]
			record[5] = v1.hours[""][3]
			record[6] = v1.hours[""][4]
			record[7] = v1.hours[""][5]
			record[8] = v1.hours[""][6]
			record[9] = v1.hours[""][7]
			record[10] = v1.hours[""][8]
			record[11] = v1.hours[""][9]
			record[12] = v1.hours[""][10]
			record[13] = v1.hours[""][11]
			record[14] = v1.costs[0]
			record[15] = v1.costs[1]
			record[16] = v1.costs[2]
			record[17] = v1.costs[3]
			record[18] = v1.costs[4]
			record[19] = v1.costs[5]
			record[20] = v1.costs[6]
			record[21] = v1.costs[7]
			record[22] = v1.costs[8]
			record[23] = v1.costs[9]
			record[24] = v1.costs[10]
			record[25] = v1.costs[11]
		}
	}
	return rawData, nil
}
