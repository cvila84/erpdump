package main

import (
	"github.com/cvila84/erpdump/internal/ebs"
)

func main() {
	//	if err := ebs.GenerateFromEBSExport("./erp2022.csv", "./erp2022-pivot.csv"); err != nil {
	//		panic(err)
	//	}
	if err := ebs.GenerateFromFinanceExport("./budget2022.csv", "./budget2022-pivot.csv"); err != nil {
		panic(err)
	}
}
