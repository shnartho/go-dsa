package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// lets create a linked list
	// linked list: 1 -> 2 -> 3 -> 4
	head := &Node{Value: 1}
	current := head
	for i := 2; i <= 5; i++ {
		newNode := &Node{Value: i}
		current.Next = newNode
		current = newNode
	}

	fmt.Println("Original Linked List:")
	PrintList(head)

	reversedHead := ReverseLinkedList(head)
	fmt.Println("Reversed Linked list:")
	PrintList(reversedHead)

}
