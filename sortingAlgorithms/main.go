package main

import (
	"example.com/user/sortingAlgorithms/abstractclasses"
	"example.com/user/sortingAlgorithms/implementation"
	"example.com/user/sortingAlgorithms/interfaces"
)

func main() {
	var comparator interfaces.Comparator
	var sorter abstractclasses.Sorter
	var processor *implementation.RecordProcessor

	comparator = &implementation.IntegerComparator{}
	sorter = implementation.SetSelectSorter(&comparator)

	processor = implementation.SetRecordProcessor(sorter, "gpa")
	processor.DoWork()
}
