package structure

import (
	"errors"
	"sort"
)

type OrderedList struct {
	data []int
}

func (ol *OrderedList) Pop() (int, error) {
	if len(ol.data) == 0 {
		return 0, errors.New("Empty data")
	}
	index := ol.data[0]
	ol.data = ol.data[1:]
	return index, nil
}

func (ol *OrderedList) Push(index int) {
	ol.data = append(ol.data, index)
	sort.Ints(ol.data)
}
