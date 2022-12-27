package main

import (
	"encoding/csv"
	"fmt"
	"github.com/cvila84/erpdump/internal/erp"
	"github.com/cvila84/erpdump/pkg/utils"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	// records[0]=manager
	// records[1]=employee
	// records[6]=hours
	// records[9]=project
	// records[10]=task
	// records[12-17]=times
	erpTimeCards := readCsvFile("./erp2022.csv")
	employeesTimes := &utils.Vector[erp.EmployeeTimes]{ID: func(element erp.EmployeeTimes) string { return element.Name }}
	for i, card := range erpTimeCards {
		if i > 0 {
			employeeTime, ok := employeesTimes.Get(card[1])
			if !ok {
				employeeTime = erp.EmployeeTimes{Name: card[1], ManagerName: card[0]}
				employeesTimes.Add(employeeTime)
			}
			month, hours1, hours2, err := erp.MonthlyHours(card)
			if err != nil {
				panic(err)
			}
			//			fmt.Printf("Record %d | %s | %s | %s | %s | %v | %d | %.2f | %.2f\n", i, record[1], record[9], record[10], record[6], hours, month, hours1, hours2)
			employeeTime.Add(erp.ParseProjectID(card[9]), card[10], month, hours1, hours2)
		}
	}
	// records[0]=employee
	// records[1]=manager
	// records[2]=project
	// records[3]=task
	// records[4]=times
	var rawData [][]interface{}
	for _, data := range employeesTimes.GetAll() {
		rawData = append(rawData, data.GetAll()...)
	}
	table := utils.NewTable(rawData).Rows(0, func(data interface{}) bool { return true })
	fmt.Println(table)
}
