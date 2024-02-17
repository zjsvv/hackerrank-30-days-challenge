package main

import "fmt"

/*
Sample Input
7
3
5
2
1
4
6
7

Sample Output
3
*/

type Node struct {
	data  int
	right *Node
	left  *Node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	} else {
		if data < root.data {
			curr := insert(root.left, data)
			root.left = curr
		} else {
			curr := insert(root.right, data)
			root.right = curr
		}
		return root
	}
}

func getHeight(root *Node) int {
	if root == nil {
		return -1
	}

	return 1 + max(getHeight(root.left), getHeight(root.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

	fmt.Println(getHeight(root))
}
