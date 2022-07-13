package main

import (
	"fmt"
	"lwj/src/command"
	"lwj/src/handle"
)

func main() {
	var templatePath = "./lwjCliTemporary"
	input := handle.GetInput()

	if input.InputTextLength == 0 {
		return
	}

	command.RunCommand("bash", "-c", "rm -rf ./lwjCliTemporary")

	command.RunCommand("bash", "-c", "git clone git@gitee.com:liweijia/lwj-project-template.git ./lwjCliTemporary")

	handle.Copy(templatePath, input.InputText)

	fmt.Println("生成完成✅️")

	command.RunCommand("bash", "-c", "rm -rf ./lwjCliTemporary")
}
