package main

import (
	"fmt"
	"github.com/cvila84/erpdump/internal/ebs"
)

func main() {
	if err := ebs.GenerateFromEBSExport([]string{"./erp-gorse-2022.csv", "./erp-grellier-2022.csv"}, "./", "pivot-erp-2022"); err != nil {
		fmt.Printf("cannot generate pivot from ebs export: %s", err)
	}
	if err := ebs.GenerateFromFinanceExport([]string{"./finance-rd-2022.csv", "./finance-l3-2022.csv"}, "./", "pivot-finance-2022"); err != nil {
		fmt.Printf("cannot generate pivot from finance export: %s", err)
	}
}
