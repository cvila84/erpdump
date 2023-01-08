package ebs

import (
	"encoding/csv"
	"github.com/cvila84/erpdump/pkg/utils"
	"os"
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

// GroupEBSTimeCardsByMonth
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
func GroupEBSTimeCardsByMonth(csvData [][]string) ([][]interface{}, error) {
	employeesTimes := &utils.Vector[EmployeeTimes]{ID: func(element EmployeeTimes) string { return element.Name }}
	for i, record := range csvData {
		if i > 0 {
			manager := strings.TrimSpace(record[0])
			employee := strings.TrimSpace(record[1])
			employeeTimes, ok := employeesTimes.Get(employee)
			if !ok {
				employeeTimes = &EmployeeTimes{Name: employee, ManagerName: manager}
				employeesTimes.Add(employeeTimes)
			}
			month, hours1, hours2, err := monthlyHours(record)
			if err != nil {
				return nil, err
			}
			if hours1 > 0 || hours2 > 0 {
				employeeTimes.Add(ParseProjectID(record[9]), record[10], month, hours1, hours2)
			}
		}
	}
	var rawData [][]interface{}
	for _, data := range employeesTimes.GetAll() {
		rawData = append(rawData, data.GetAll()...)
	}
	return rawData, nil
}

// FilterBudgetPivotData
// record[3]=month (yyyy-mm)
// record[14]=project
// record[21]=cost
// record[28]=type
// record[32]=employee
// record[40]=hours
// -->
// record[0]=type
// record[1]=employee
// record[2]=project
// record[3-14]=hours
// record[15-26]=cost

func FilterBudgetPivotData(csvData [][]string) ([][]interface{}, error) {
	var result [][]interface{}
	//for i, record := range csvData {
	//	if i > 0 {
	//
	//	}
	//}
	return result, nil
}
