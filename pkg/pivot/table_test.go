package pivot

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	if parentHeaderLabel("") != "" {
		t.Fatalf("parentHeaderLabel(\"\")=%s!=\"\"", parentHeaderLabel(""))
	}
	if parentHeaderLabel("A1") != "" {
		t.Fatalf("parentHeaderLabel(\"A1\")=%s!=\"\"", parentHeaderLabel("A1"))
	}
	if parentHeaderLabel("A1/B1") != "A1" {
		t.Fatalf("parentHeaderLabel(\"A1/B1\")=%s!=\"A1\"", parentHeaderLabel("A1/B1"))
	}
	if parentHeaderLabel("A1/B1/C1") != "A1/B1" {
		t.Fatalf("parentHeaderLabel(\"A1/B1/C1\")=%s!=\"A1/B1\"", parentHeaderLabel("A1/B1/C1"))
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
	table := NewTable(rawData, false).
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
	table = NewTable(rawData, false).
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

func TestComputeSet(t *testing.T) {
	rawData := [][]interface{}{
		{"A", "B", "V1", "V2"},
		{"A1", "B1", 4, 6},
		{"A1", "B2", 2, 4},
		{"A1", "B1", 3, 5},
	}
	compute := func(elements []interface{}) (float64, error) {
		return (elements[0].(float64)) / (elements[1].(float64)), nil
	}
	table := NewTable(rawData, true).
		StandardRow(0).
		StandardColumn(1).
		StandardValues(2, Sum).
		StandardValues(3, Sum).
		Values([]int{2, 3}, compute, Set, Digits(2))
	err := table.Generate()
	if err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(table.ToCSV())
}
