package main

import (
	"fmt"
	"os/exec"
)

func main() {

	program := "/c/Program\\ Files/Git/cmd/git.exe" //os.Args[0]
	args := "status"                                //os.Args[1:]

	fmt.Println(program)
	fmt.Println(args)

	cmd := exec.Command(program, args)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Waiting for command to finish...")
	err = cmd.Wait()
	fmt.Println("Command finished with error: %v", err)
}
