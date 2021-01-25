package entities

// Record representing data strucutres
type Record struct {
	name          string
	studentNumber int
	gpa           float32
	sortKeyName   string
}

//SetSortKeyName method to set the sortKeyName field of
// instantiated Record
func (record *Record) SetSortKeyName(keyName string) {
	(*record).sortKeyName = keyName
}

//GetSortKeyName method to get the sortKeyName field of
// instantiated Record
func (record *Record) GetSortKeyName() string {
	return record.sortKeyName
}

//StringRepr method to get the string representation of
// Record instance fields
func (record *Record) StringRepr() string {
	return "hello from record object"
}

//SetRecord function acting as constructor of record struct
// to instantiate the private variables and return the
// pointer to the instantiated record object.
func SetRecord(studentNumber int, name string, gpa float32) *Record {
	return &Record{name: name, studentNumber: studentNumber, gpa: gpa}
}
