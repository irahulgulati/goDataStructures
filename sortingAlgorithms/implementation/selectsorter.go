package implementation

import (
	"fmt"

	"example.com/user/sortingAlgorithms/abstractclasses"
	"example.com/user/sortingAlgorithms/entities"
	"example.com/user/sortingAlgorithms/interfaces"
)

// SelectSorter struct implementing
// Sorter interface
type SelectSorter struct {
	abstractclasses.AbstractSorter
}

// Sort method implementing the
// Sort method of Sorter inteface
func (s *SelectSorter) Sort(records []*entities.Record) {
	// rr := *&s.AbstractSorter
	// abc := *&rr
	// def := *abc.Comparator
	// result := def.Compare(1, 2)
	// fmt.Println(result)
	comparator := *(&(*(&(*s).AbstractSorter))).Comparator
	fmt.Println(comparator.Compare(1, 2))
}

// SetSelectSorter function acting as constructor to
// instantiate SelectSorter object and return pointer
// to it
func SetSelectSorter(comparator *interfaces.Comparator) *SelectSorter {
	superSorter := abstractclasses.AbstractSorter{
		Comparator: comparator,
	}
	var SelectSorter SelectSorter = SelectSorter{
		AbstractSorter: superSorter,
	}
	return &SelectSorter

}
