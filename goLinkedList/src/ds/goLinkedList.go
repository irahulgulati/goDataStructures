package ds

// LinkedList function to create an instance of singlyLinkedList

// node structure defining the single node of the LinkedList
type node struct {
	Next *node
	Data int
}

// SinglyLinkedList struct representing the singlylinkedclass
type SinglyLinkedList struct {
	Head *node
}

// AddNode method linked to SinglyLinkedList struct used for adding
// a node to to the singly linked list instance
func (l *SinglyLinkedList) AddNode(data int) {
	if l.Head == nil {
		newNode := &node{
			Data: data,
		}
		l.Head = newNode
		return
	}

	var traversingPtr *node
	traversingPtr = l.Head

	for traversingPtr.Next != nil {
		traversingPtr = traversingPtr.Next
	}
	traversingPtr.Next = &node{
		Data: data,
	}
	return
}
