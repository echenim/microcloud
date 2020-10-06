package main

//26028f81c5ec522c80b21d909a2333b6357cf74f

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/echenim/microcloud/utils"
)

// container entry point
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what?S???")

	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	utils.Check(cmd.Run(), "FAILED")
}
