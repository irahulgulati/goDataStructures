package main

import (
	"ds"
	"fmt"
	"os"
)

func main() {
	var singlyLinkedList ds.SinglyLinkedList
	var error error

	singlyLinkedList = ds.SinglyLinkedList{}
	(&singlyLinkedList).AddNode(30)
	(&singlyLinkedList).AddNode(20)
	fmt.Println(singlyLinkedList.Head.Next.Data)
	if error != nil {
		os.Exit(1)
	}
}
