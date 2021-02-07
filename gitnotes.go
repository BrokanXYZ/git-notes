package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// ForwardGitCommand ...
func ForwardGitCommand(args []string) {

	program := "git"
	cmd := exec.Command(program, args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// GitToMarkdown ...
func GitToMarkdown(args []string) {
	fmt.Println(args[0])
	switch command := args[0]; command {
	case "commit":
		fmt.Println("COMMIT")
	case "checkout":
	case "branch":
		fmt.Println("NEW BRANCH??")
	}
}

func main() {
	args := os.Args[1:]
	GitToMarkdown(args)
	ForwardGitCommand(args)
}
