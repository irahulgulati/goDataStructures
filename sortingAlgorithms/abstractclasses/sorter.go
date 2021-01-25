package abstractclasses

import (
	"example.com/user/sortingAlgorithms/entities"
	"example.com/user/sortingAlgorithms/interfaces"
)

// Sorter interface that will be
// extended by SelectSorter and
// BubbleSorter
type Sorter interface {
	Sort(records []*entities.Record)
}

type AbstractSorter struct {
	Comparator *interfaces.Comparator
}
