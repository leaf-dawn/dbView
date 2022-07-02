package tree

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	root := Deserialize("{1,2,3,4,5,6,6,7,8,9,10,11,12,#,#,#,#,#,#,#,#,#,#,#,#,#,#,#,#}")
	treePrintBuf := NewTreePrintBuf()
	treePrintBuf.PutTree(root, root)
	treePrintBuf.Print()

}

func Test2(t *testing.T) {
	treePrintBuf := NewTreePrintBuf()
	root := Deserialize("{1,2,3,4,5,6,6,#,#,#,#,#,#,#,#}")
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		treePrintBuf.PutTree(root, node)
		stack = stack[0 : len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		treePrintBuf.Print()
		time.Sleep(time.Second * 2)
		treePrintBuf.Clear()
	}
}
