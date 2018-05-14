package cmd

import (
	"os/exec"
	"log"
)

func ExecCmd(command string) error {
	cmd := exec.Command("cmd.exe", "/c", command)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("命令 \"%s\" 执行完毕.\n", "cmd.exe /c "+command)
	return nil
}
