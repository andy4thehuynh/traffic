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
	if err != nil {
		log.Fatalln(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Start()

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	cmd.Wait()

	if cmd.ProcessState.Success() {
		fmt.Println("It finished bro")
		SetTmuxStatusColor("green")
	} else {
		fmt.Println("naw you didn't success it")
		SetTmuxStatusColor("red")
	}
}

func SetTmuxStatusColor(color string) error {
	cmd := exec.Command("tmux", "set", "status-bg", color)
	return cmd.Run()
}
