package table

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	if parentLabel("") != "" {
		t.Fatalf("parentLabel(\"\")=%s!=\"\"", parentLabel(""))
	}
	if parentLabel("A1") != "" {
		t.Fatalf("parentLabel(\"A1\")=%s!=\"\"", parentLabel("A1"))
	}
	if parentLabel("A1/B1") != "A1" {
		t.Fatalf("parentLabel(\"A1/B1\")=%s!=\"A1\"", parentLabel("A1/B1"))
	}
	if parentLabel("A1/B1/C1") != "A1/B1" {
		t.Fatalf("parentLabel(\"A1/B1/C1\")=%s!=\"A1/B1\"", parentLabel("A1/B1/C1"))
	}
	rawData := [][]interface{}{
		{"A1", "B1", "C1", "D1", 4},
		{"A1", "B2", "C1", "D1", 2},
		{"A1", "B1", "C2", "D1", 3},
		{"A1", "B1", "C2", "D2", 1},
		{"A2", "B1", "C1", "D2", 5},
		{"A1", "B1", "C2", "D1", 1},
	}
	//           D1      D2      Total
	// A1        10      1       11
	// A1/B1     8       1       9
	// A1/B1/C1  4               4
	// A1/B1/C2  4       1       5
	// A1/B2     2               2
	// A1/B2/C1  2               2
	// A2                5       5
	// A2/B1             5       5
	// A2/B1/C1          5       5
	// Total     10      6       16
	table := NewIntTable(rawData).
		StandardRow(0).
		StandardRow(1).
		StandardRow(2).
		StandardColumn(3).
		StandardValues(4, Sum)
	err := table.Generate()
	if err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(table.ToCSV())
	table = NewIntTable(rawData).
		StandardRow(0).
		StandardRow(1).
		StandardColumn(2).
		StandardColumn(3).
		StandardValues(4, Sum)
	err = table.Generate()
	if err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(table.ToCSV())
}
