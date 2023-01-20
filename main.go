package main

import (
	"github.com/cvila84/erpdump/internal/ebs"
)

func main() {
	if err := ebs.GenerateFromEBSExport("./erp-2022.csv", "./", "erp-2022-pivot"); err != nil {
		panic(err)
	}
	if err := ebs.GenerateFromFinanceExport([]string{"./budget-rd-2022.csv", "./budget-l3-2022.csv"}, "./", "pivot-budget-2022"); err != nil {
		panic(err)
	}
}
