package _map

import (
	"strconv"
	"strings"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func Deserialize(s string) *Node {
	var beginNode *Node
	m := map[int]*Node{}
	//转化为数组
	slist := strings.Split(s[2:len(s)-2], "],[")
	for i, v := range slist {
		//判断i节点是否在m中
		if _, ok := m[i+1]; !ok {
			m[i+1] = &Node{Val: i + 1, Neighbors: []*Node{}}
		}
		//添加边
		list := strings.Split(v, ",")
		for _, v2 := range list {
			num, _ := strconv.Atoi(v2)
			if _, ok := m[num]; !ok {
				m[num] = &Node{Val: num, Neighbors: []*Node{}}
			}
			m[i+1].Neighbors = append(m[i+1].Neighbors, m[num])
		}
		if beginNode == nil {
			beginNode = m[i]
		}
	}
	return beginNode
}
