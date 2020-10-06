package namespaces

import (
	"os/exec"
	"syscall"
)

//GetProcessID function for getting the process id
func (n *Namespace) GetProcessID(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}
}
