package command

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	log.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
