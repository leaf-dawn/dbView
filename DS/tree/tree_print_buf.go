package tree

import (
	"fmt"
	"math"
	"strconv"
	print2 "test/DS/print"
)

/**
 * 树相关输出方法
 * 用来动态输出树结构的
 */

type TreePrintBuf interface {
	PutTree(root *TreeNode, node *TreeNode)
	Print()
	ScanToClear()
	Clear()
}

type treePrintBuf struct {
	terminal *print2.SimpleTerminal
}

func NewTreePrintBuf() *treePrintBuf {
	return &treePrintBuf{
		terminal: print2.NewTerminal(),
	}
}

//打印一颗树，node表示当前节点
func (tm *treePrintBuf) PutTree(root *TreeNode, highLightNode *TreeNode) {
	depth, bigSize := getDepthAndNodeSize(root)
	if depth == 0 {
		return
	}
	mid := int(math.Pow(2, float64(depth)-1)) * (bigSize + 2) / 2
	//每一层x的起始位置
	beginX := mid - (bigSize+2)/2

	nodeQueue := []*TreeNode{root}
	nodeY := 0 //记录y的位置
	deep := 1
	for len(nodeQueue) != 0 {
		//到这一层的起始位置
		tm.terminal.GotoXY(beginX, nodeY)
		l := len(nodeQueue)
		for i := 0; i < l; i++ {
			node := nodeQueue[0]
			nodeQueue = nodeQueue[1:]
			//添加node节点，并到下一个位置
			tm.nodeAdd(bigSize+2, node, node == highLightNode)
			if node != nil {
				nodeQueue = append(nodeQueue, node.Left)
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		//更新到下一层
		beginX = mid - int(math.Pow(2, float64(deep)-1))*(bigSize+2)
		nodeY += 2
		//记录深度
		deep++
	}
}

func (tm *treePrintBuf) Print() {
	tm.terminal.Print()
}

func (tm *treePrintBuf) Clear() {
	tm.terminal.Clear()
}

func (tm *treePrintBuf) ScanToClear() {
	scan := 0
	_, _ = fmt.Scan(&scan)
	tm.terminal.Clear()
}

//获取树的深度，和最大输出节点值的长度
func getDepthAndNodeSize(root *TreeNode) (int, int) {
	bigSize := 0
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) != 0 {
		depth++
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Val/10+1 > bigSize {
				bigSize = node.Val/10 + 1
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth, bigSize
}

//打印一个节点，并添加下一个节点位置到切片中
func (sm treePrintBuf) nodeAdd(nodeSize int, node *TreeNode, isHighLight bool) {
	x := sm.terminal.GetX()
	y := sm.terminal.GetY()
	if node != nil {
		//打印普通节点
		buf := make([]byte, nodeSize)
		copy(buf[1:], strconv.Itoa(node.Val))
		if isHighLight {
			sm.terminal.AddAndType(string(buf),
				print2.GetFormat(print2.DISPLAY_HIGH_LIGHT, print2.PROSPECT_RED))
		} else {
			sm.terminal.Add(string(buf))
		}
		sm.terminal.GotoXY(x, sm.terminal.GetY()+1)
		buf = make([]byte, nodeSize)
		buf[0] = '/'
		buf[len(buf)-1] = '\\'
		sm.terminal.Add(string(buf))
	}
	//走到下一个位置
	sm.terminal.GotoXY(x+nodeSize, y)
}
