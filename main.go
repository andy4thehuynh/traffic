package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	SetTmuxStatusColor("yellow")
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Not enough arguments")
	}

	cmd := exec.Command(args[0], args[1:]...)

	stdout, err := cmd.StdoutPipe()
	ExitIfErr(err)

	stderr, err := cmd.StderrPipe()
	ExitIfErr(err)

	ExitIfErr(cmd.Start())

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	ExitIfErr(cmd.Wait())

	if cmd.ProcessState.Success() {
		SetTmuxStatusColor("green")
	} else {
		SetTmuxStatusColor("red")
	}
}

func SetTmuxStatusColor(color string) error {
	cmd := exec.Command("tmux", "set", "status-bg", color)
	return cmd.Run()
}

func ExitIfErr(err error) {
	if err != nil {
		Exit(err)
	}
}

func Exit(message interface{}) {
	SetTmuxStatusColor("red")
	fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
