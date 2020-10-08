package main

import (
	"os"

	"github.com/echenim/microcloud/core"
)

// container entry point
func main() {

	switch os.Args[1] {
	case "run":
		core.Parent()
	case "child":
		core.Child()
	default:
		panic("what??")

	}

}
