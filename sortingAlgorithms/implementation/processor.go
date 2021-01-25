package implementation

import (
	"fmt"

	"example.com/user/sortingAlgorithms/abstractclasses"
	"example.com/user/sortingAlgorithms/entities"
)

//RecordProcessor struct as a processor
type RecordProcessor struct {
	sorter         abstractclasses.Sorter
	sortingKeyName string
}

//DoWork method to start the whole sorting
// process
func (r *RecordProcessor) DoWork() {
	fmt.Println("doing work")
	takeInputs()
	var rec *entities.Record
	rec = entities.SetRecord(
		1,
		"Rahul",
		4.1,
	)
	rec.SetSortKeyName("gpa")
	fmt.Println(rec)
	inputRecords := []*entities.Record{
		rec,
	}
	(*r).sorter.Sort(inputRecords)
}

func takeInputs() {
	fmt.Println("taking inputs")
}

func showOutput() {
	fmt.Println("Showing output")
}

//SetRecordProcessor construtor that instantiate
// RecordProcessor struct and return pointer to it
func SetRecordProcessor(sorter abstractclasses.Sorter, sortingKeyName string) *RecordProcessor {
	return &RecordProcessor{sorter: sorter, sortingKeyName: sortingKeyName}
}
