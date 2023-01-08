package main

import (
	"github.com/cvila84/erpdump/internal/ebs"
)

func main() {
	if err := ebs.GenerateFromEBSExport("./erp2022.csv", "./erp2022-pivot.csv"); err != nil {
		panic(err)
	}
}
