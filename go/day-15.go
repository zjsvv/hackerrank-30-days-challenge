package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func insert(head *Node, data int) *Node {
	if head == nil {
		head = &Node{Data: data}
	} else {
		curr := head
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = &Node{Data: data}
	}

	return head
}

func display(head *Node) {
	start := head

	for start != nil {
		fmt.Printf("%d ", start.Data)
		start = start.Next
	}
}

func main() {
	// scan input
	var n int
	fmt.Scan(&n)

	var head *Node

	for n > 0 {
		var data int
		fmt.Scan(&data)
		head = insert(head, data)
		n--
	}

	display(head)
}
