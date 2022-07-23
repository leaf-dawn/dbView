package print

import (
	"fmt"
)

/**
 * 终端输出工具
 * 虚假的控制光标,仅仅是使用buf进行存储然后输出
 */

/**
 *设置终端
 */
type Terminal interface {
	//获取当前x位置
	GetX() int
	//获取当前y位置
	GetY() int
	//把光标移动到某个位置
	GotoXY(x int, y int)
	//光标当前位置向x,y移动x，y的距离
	Move(lenX int, lenY int)
	//输出
	Add(str string)
	//输出+类型
	AddAndType(str string, tye string)
	//换行
	NewLine()
	//打印
	Print()
	//清除
	Clear()
}

type SimpleTerminal struct {
	x            int
	y            int
	terminalByte [][]byte
	types        [][]string
}

func NewTerminal() *SimpleTerminal {
	return &SimpleTerminal{
		x:            0,
		y:            0,
		terminalByte: make([][]byte, 1, 10),
		types:        make([][]string, 1, 10),
	}
}

func (st *SimpleTerminal) GotoXY(x int, y int) {
	st.expandBuf(x+1, y+1)
	st.x = x
	st.y = y
}

func (st *SimpleTerminal) Move(lenX int, lenY int) {
	st.expandBuf(st.x+lenX+1, st.y+lenY+1)
	st.x = st.x + lenX
	st.y = st.y + lenY
}

func (st *SimpleTerminal) Add(str string) {
	st.expandBuf(st.x+len(str)+1, st.y+1)
	strBuf := []byte(str)
	for i := 0; i < len(strBuf); i++ {
		st.terminalByte[st.y][st.x] = strBuf[i]
		st.x++
	}
}

func (st *SimpleTerminal) AddAndType(str string, tye string) {
	st.expandBuf(st.x+len(str)+1, st.y+1)
	strBuf := []byte(str)
	for i := 0; i < len(strBuf); i++ {
		st.terminalByte[st.y][st.x] = strBuf[i]
		st.types[st.y][st.x] = tye
		st.x++
	}
}

func (st *SimpleTerminal) NewLine() {
	st.y++
	st.x = 0
	st.expandBuf(st.x, st.y)
}

func (st *SimpleTerminal) Print() {
	//遍历所有的并输出
	for i := 0; i < len(st.terminalByte); i++ {
		for j := 0; j < len(st.terminalByte[i]); j++ {
			if st.types[i][j] == "" {
				fmt.Print(string(st.terminalByte[i][j]))
			} else {
				fmt.Printf(st.types[i][j], string(st.terminalByte[i][j]))
			}
		}
		fmt.Println()
	}
}

func (st *SimpleTerminal) GetX() int {
	return st.x
}

func (st *SimpleTerminal) GetY() int {
	return st.y
}

func (st *SimpleTerminal) Clear() {
	Clear()
	st.terminalByte = make([][]byte, 1, 10)
	st.types = make([][]string, 1, 10)
	st.x = 0
	st.y = 0
}

/**
 * len表示需要扩充的最小长度
 * 扩大切片
 */
func (st *SimpleTerminal) expandBuf(lenX int, lenY int) {
	//扩容y
	if lenY > len(st.terminalByte) {
		l := len(st.terminalByte)
		for i := 0; i < lenY-l; i++ {
			st.terminalByte = append(st.terminalByte, make([]byte, 10))
			st.types = append(st.types, make([]string, 10))
		}
	}
	//扩容x
	if lenX > len(st.terminalByte[lenY-1]) {
		newBufX := make([]byte, 2*lenX)
		copy(newBufX, st.terminalByte[lenY-1])
		st.terminalByte[lenY-1] = newBufX
		newBufStr := make([]string, 2*lenX)
		copy(newBufStr, st.types[lenY-1])
		st.types[lenY-1] = newBufStr
	}
}
