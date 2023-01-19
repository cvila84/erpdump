package pivot

import (
	"fmt"
	"strings"
)

type Operation int

const (
	none Operation = iota
	Count
	Sum
)

type cellRecordKey struct {
	dataIndex int
	operation Operation
}

type cell[T valueType] interface {
	fmt.Stringer
	Compute(index int, compute Compute[T], keys []cellRecordKey) error
	Update(index int, key cellRecordKey)
	Get() []T
	Record(key cellRecordKey, value T)
}

type pivotCell[T valueType] struct {
	finalValues    []T
	recordedValues map[cellRecordKey]T
	formats        []string
}

func newPivotCell(formats []ValueFormat) cell[float64] {
	valueFormats := make([]string, len(formats))
	for i := 0; i < len(formats); i++ {
		valueFormats[i] = string(formats[i])
	}
	return &pivotCell[float64]{
		finalValues:    make([]float64, len(formats)),
		recordedValues: make(map[cellRecordKey]float64),
		formats:        valueFormats,
	}
}

func (p *pivotCell[T]) String() string {
	var sb strings.Builder
	if len(p.finalValues) > 1 {
		sb.WriteString("[ ")
		for i := 0; i < len(p.finalValues); i++ {
			sb.WriteString(fmt.Sprintf(p.formats[i], p.finalValues[i]))
			if i < len(p.finalValues)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(" ]")
		return sb.String()
	} else {
		return fmt.Sprintf(p.formats[0], p.finalValues[0])
	}
}

func (p *pivotCell[T]) Compute(index int, compute Compute[T], keys []cellRecordKey) error {
	var elements []RawValue
	var err error
	for _, k := range keys {
		elements = append(elements, p.recordedValues[k])
	}
	p.finalValues[index], err = compute(elements)
	return err
}

func (p *pivotCell[T]) Update(index int, key cellRecordKey) {
	p.finalValues[index] += p.recordedValues[key]
}

func (p *pivotCell[T]) Get() []T {
	return p.finalValues
}

func (p *pivotCell[T]) Record(key cellRecordKey, value T) {
	if key.operation == Sum {
		p.recordedValues[key] += value
	} else if key.operation == Count {
		p.recordedValues[key]++
	}
}
