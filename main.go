package main

import (
	"github.com/cvila84/erpdump/internal/ebs"
)

func main() {
	if err := ebs.GenerateFromEBSExport("./erp-2022.csv", "./erp-2022-pivot.csv"); err != nil {
		panic(err)
	}
	if err := ebs.GenerateFromFinanceExport("./budget-2022.csv", "./budget-2022-pivot.csv"); err != nil {
		panic(err)
	}
	if err := ebs.GenerateFromFinanceExport("./budgetl3-2022.csv", "./budgetl3-2022-pivot.csv"); err != nil {
		panic(err)
	}
}
