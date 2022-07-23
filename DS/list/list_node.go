package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLink(head *ListNode) {

	i := 0
	fmt.Print("[")
	for head != nil && i < 10000 {
		i++
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Print("]")
}
