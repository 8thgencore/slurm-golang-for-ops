package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(slice []int) *ListNode {
	node := &ListNode{
		Val:  slice[len(slice)-1],
		Next: nil,
	}

	for i := len(slice) - 2; i >= 0; i-- {
		node = &ListNode{
			Val:  slice[i],
			Next: node,
		}
	}

	return node
}

func deleteDuplicates(n *ListNode) {
	for n.Next != nil {
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
}

func nodesPrint(n ListNode) {
	for n.Next != nil {
		fmt.Printf("List nodes start: %v\n", n)
		n = *n.Next
	}
}

func main() {
	slice := []int{13, 17, 23, 45, 45, 57, 59}
	nodes := createList(slice)

	fmt.Printf("=== List nodes start ===\n")
	nodesPrint(*nodes)

	deleteDuplicates(nodes)
	fmt.Printf("\n=== List nodes after clean ===\n")
	nodesPrint(*nodes)
}
