package print

import (
	"strconv"
)

/**
 * 字体相关工具
 */

const (
	//显示格式
	DISPLAY_HIGH_LIGHT  = 1
	DISPLAY_DEFAULT     = 0
	DISPLAY_UNDER_SCORE = 4
	DISPLAY_FLICKER     = 5
	//背景颜色
	BACK_BLACK  = 40
	BACK_RED    = 41
	BACK_GREEN  = 42
	BACK_YELLOW = 43
	BACK_BLUE   = 44
	BACK_WHITE  = 47
	//字体颜色
	PROSPECT_BLACK  = 30
	PROSPECT_RED    = 31
	PROSPECT_GREEN  = 32
	PROSPECT_YELLOW = 33
	PROSPECT_BLUE   = 34
	PROSPECT_WHITE  = 37
)

func GetFormat(formatCode ...int) string {
	beginStr := "\033["
	for i := 0; i < len(formatCode); i++ {
		beginStr = beginStr + strconv.Itoa(formatCode[i]) + ";"
	}
	beginStr = beginStr[:len(beginStr)-1] + "m"
	return beginStr + "%s\033[0m"
}
