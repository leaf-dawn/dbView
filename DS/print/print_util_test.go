package print

import (
	"testing"
)

func TestTerminal(t *testing.T) {
	terminal := NewTerminal()
	terminal.Add("232")
	terminal.Move(7, 0)
	terminal.Add("234")
	terminal.NewLine()
	terminal.Add("234")
	terminal.GotoXY(3, 7)
	terminal.Add("goto")
	terminal.AddAndType("color", GetFormat(DISPLAY_HIGH_LIGHT, BACK_BLUE, PROSPECT_RED))
	terminal.Print()
}
