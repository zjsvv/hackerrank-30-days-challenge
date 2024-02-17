package main

import "fmt"

/*
Sample Input
6
3
5
4
7
2
1

Sample Output
3 2 5 1 4 7
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

func levelOrder(root *Node) {
	if root == nil {
		return
	}

	q := make([]*Node, 0)

	q = append(q, root)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		fmt.Printf("%d ", node.data)

		if node.left != nil {
			q = append(q, node.left)
		}

		if node.right != nil {
			q = append(q, node.right)
		}
	}
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

	levelOrder(root)
}
