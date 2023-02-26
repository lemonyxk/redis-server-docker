/**
* @program: redis-sentinel-docker
*
* @description:
*
* @author: lemo
*
* @create: 2023-02-26 15:37
**/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func GetArgs(flag ...string) string {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				if i+1 < len(args) {
					return args[i+1]
				}
			}
		}
	}
	return ""
}

func HasArgs(flag ...string) bool {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				return true
			}
		}
	}
	return false
}

func Command(command ...string) *exec.Cmd {
	if len(command) == 0 {
		exitWithMessage("command is empty")
	}
	var cmd = exec.Command(command[0], command[1:]...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return cmd
}

func exitWithMessage(msg interface{}) {
	fmt.Printf("%+v\n", msg)
	os.Exit(0)
}
