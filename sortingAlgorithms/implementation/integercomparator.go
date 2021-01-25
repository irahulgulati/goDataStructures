package implementation

//IntegerComparator struct implementing
// comparator abstract class
type IntegerComparator struct{}

// Compare method implementing compare method of
// abstract class
func (i IntegerComparator) Compare(a int, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}
