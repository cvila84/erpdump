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
		Row(0).
		Row(1).
		Row(2).
		Column(3).
		Values(4, Sum, Digits(0))
	err := table.Generate()
	if err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(table.ToCSV())
	table = NewTable(rawData, false).
		Row(0).
		Row(1).
		Column(2).
		Column(3).
		Values(4, Sum, Digits(0))
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
		{"A1", "B1", 8, 5},
		{"A1", "B2", 4, 2},
	}
	compute := func(elements []interface{}) (float64, error) {
		return (elements[0].(float64)) / (elements[1].(float64)), nil
	}
	table := NewTable(rawData, true).
		Row(0).
		Column(1).
		ComputedValues([]int{3, 2}, compute, Digits(2)).
		Values(3, Sum, Digits(0))
	err := table.Generate()
	if err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(table.ToCSV())
}

//pivot         map[string]map[string][]cell
//rowHeaders    *headers
//columnHeaders *headers
//rowSeries     []*series[string]
//columnSeries  []*series[string]
//valueSeries   []*series[float64]

//indexes       map[int]int
// 0 => 0
// 1 => 0
// 2 => 0
// 3 => 1
