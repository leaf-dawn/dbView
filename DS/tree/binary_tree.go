package tree

import (
	"strconv"
	"strings"
)

//普通二叉树
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Serialize(root *BinaryTreeNode) string {
	// write code here
	//层序遍历
	if root == nil {
		return "{}"
	}
	queue := []*BinaryTreeNode{root}
	str := "{"
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			str = str + strconv.Itoa(node.Val) + ","
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			str = str + "#,"
		}
	}
	return str[:len(str)-1] + "}"
}

//反序列化树
func Deserialize(s string) *BinaryTreeNode {
	// write code here
	//先序遍历
	//添加第一个元素到root中
	s = s[1 : len(s)-1]
	listStr := strings.Split(s, ",")
	if len(listStr) == 1 && listStr[0] == "" {
		return nil
	}
	rootStr := listStr[0]
	listStr = listStr[1:]
	rootVal, _ := strconv.Atoi(rootStr)
	root := &BinaryTreeNode{
		Val: rootVal,
	}
	nodeQueue := []*BinaryTreeNode{root}
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		lStr := listStr[0]
		listStr = listStr[1:]
		rStr := listStr[0]
		listStr = listStr[1:]
		if lStr == "#" {
			node.Left = nil
		} else {
			v, _ := strconv.Atoi(lStr)
			node.Left = &BinaryTreeNode{
				Val: v,
			}
			nodeQueue = append(nodeQueue, node.Left)
		}
		if rStr == "#" {
			node.Right = nil
		} else {
			v, _ := strconv.Atoi(rStr)
			node.Right = &BinaryTreeNode{
				Val: v,
			}
			nodeQueue = append(nodeQueue, node.Right)
		}

	}
	return root
}
