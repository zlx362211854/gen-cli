package handle

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)
// 获取输入
type InputResult struct  {
	InputText string // 输入值
	InputTextLength int // 值长度
}
func GetInput() InputResult {
	var text = ""
	var textLength = 0
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("️请输入项目名称:\n")
	for input.Scan() {
		line := input.Text()
		len := utf8.RuneCountInString(line)
		if len == 0 {
			fmt.Printf("⚠请输入项目名称:\n")
		}
		if line == "bye" {
			break
		}
		if len > 0 {
			fmt.Println("⏳ 正在生成项目：" + line)
			text = line
			textLength = len
			break
		}
	}
	return InputResult{
		InputText: text,
		InputTextLength: textLength,
	}
}