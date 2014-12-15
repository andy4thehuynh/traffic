package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Listen to status codes from unix commands.
// 0 - Green.. it's done
// 1 - Red.. something went wrong
// other - Yellow.. still in progress

// Change status bar color

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Not enough arguments")
	}

	cmd := exec.Command(args[0], args[1:]...)

	data, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(data))
}
