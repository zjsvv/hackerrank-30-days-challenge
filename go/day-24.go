package main

import "fmt"

/*
Sample Input
6
1
2
2
3
3
4

Sample Output
1 2 3 4
*/

type Node struct {
	data  int
	next *Node
}

func insert(head *Node, data int) *Node {
	p := &Node{data: data}

	if head == nil {
		head = p
	} else if head.next == nil {
		head.next = p
	} else {
		start := head

		for start.next != nil {
			start = start.next
		}
		start.next = p
	}
	return head
}

func display(head *Node) {
	start := head

	for start != nil {
		fmt.Printf("%d ", start.data)
		start = start.next
	}
}

func removeDuplicates(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head

	for curr != nil {
		if curr.next != nil && curr.data == curr.next.data {
			curr.next = curr.next.next
			continue
		}
		curr = curr.next
	}

	return head
}

func main() {
	var n int
	fmt.Scan(&n)

	var root *Node

	for n > 0 {
		var data int
		fmt.Scan(&data)
		root = insert(root, data)
		n--
	}

	root = removeDuplicates(root)
	display(root)
}
