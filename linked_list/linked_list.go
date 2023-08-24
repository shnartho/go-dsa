package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{Value: value, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) DeleteNode(value int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Value != value {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}

}

func (ll *LinkedList) SearchNode(value int) bool {
	current := ll.Head
	if current != nil {
		for current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList) PrintList() {
	current := ll.Head
	if current == nil {
		fmt.Println("Linke list doesn't exist")
	}
	for current != nil {
		fmt.Printf("-> %d", current.Value)
		current = current.Next
	}
}

func main() {
	ll := &LinkedList{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. Insert at beginning")
		fmt.Println("2. Insert at end")
		fmt.Println("3. Delete node")
		fmt.Println("4. Search node")
		fmt.Println("5. Print list")
		fmt.Println("6. Exit")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1, 2, 3, 4:
			fmt.Println("Please enter value")
			valueStr, _ := reader.ReadString('\n')
			valueStr = strings.TrimSpace(valueStr)
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				fmt.Println("Invalid value. Please enter a valid number")
				continue
			}
			switch choice {
			case 1:
				ll.InsertAtBeginning(value)
			case 2:
				ll.InsertAtEnd(value)
			case 3:
				ll.DeleteNode(value)
			case 4:
				found := ll.SearchNode(value)
				if found {
					fmt.Println("Value", value, " found in the linked list.")
				} else {
					fmt.Println("Value", value, " not found in the linked list")
				}
			}

		case 5:
			fmt.Println("Printing Linked List: ")
			ll.PrintList()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Sprintln("Invalid choice. Please enter a valid number.")
		}
	}

}
