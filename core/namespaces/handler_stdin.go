package namespaces

import (
	"os"
	"os/exec"
)

//GetInput from system karnel
func (n Namespace) GetInput(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
}

//GetOutput from system karnel
func (n Namespace) GetOutput(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
}

//GetError from system karnel
func (n Namespace) GetError(cmd *exec.Cmd) {
	cmd.Stderr = os.Stderr
}
