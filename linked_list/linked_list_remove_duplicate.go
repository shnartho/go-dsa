package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func RemoveDuplicates(head *Node) {
	current := head
	for current != nil && current.Next != nil {
		if current.Value == current.Next.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("-> %d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// create a sorted linked list with duplicates: 1 ->2 ->3 ->3 ->3 ->4
	head := &Node{Value: 1}
	current := head
	values := []int{2, 3, 3, 3, 4}
	for _, value := range values {
		newNode := &Node{Value: value}
		current.Next = newNode
		current = newNode
	}

	fmt.Println("Original linked list:")
	PrintLinkedList(head)

	RemoveDuplicates(head)

	fmt.Println("After removing duplicates")
	PrintLinkedList(head)

}
