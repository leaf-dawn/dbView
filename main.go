package main

import (
	"fmt"
)

type test2 struct {
	a int32 //4
	c int32 // 4
	b int64 // 8
}

/**
| a     |        |       b       |   c    |
 * * * * _ _ _ _ * * * * * * * * | * * * * _ _ _ _
*/
func main() {
	fmt.Println(findMinHeightTrees(7, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}))

}

//拓扑排序解
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	edgesl := make([][]int, n)
	degrees := make([]int, n)
	for i := range edges {
		x := edges[i][0]
		y := edges[i][1]
		edgesl[x] = append(edgesl[x], y)
		edgesl[y] = append(edgesl[y], x)
		//添加度数
		degrees[x]++
		degrees[y]++
	}
	//遍历获取所有入度为1的节点
	queue := []int{}
	for k, v := range degrees {
		if v == 1 {
			queue = append(queue, k)
		}
	}
	remain := n
	for remain > 2 {
		l := len(queue)
		remain -= l
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			//删除该节点
			for _, v := range edgesl[node] {
				degrees[v]--
				if degrees[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return queue
}
