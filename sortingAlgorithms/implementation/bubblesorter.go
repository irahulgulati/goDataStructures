package implementation

import (
	"fmt"

	"example.com/user/sortingAlgorithms/abstractclasses"
	"example.com/user/sortingAlgorithms/entities"
)

//BubbleSorter struct
type BubbleSorter struct {
	abstractclasses.Sorter
}

// Sort , BubbleSorter struct method implementing
// interface method
func (b *BubbleSorter) Sort(records []entities.Record) {
	fmt.Println("hello from sorter")
}
