package namespaces

import (
	"os/exec"
	"syscall"
)

//GetHostName function for getting the host name
//for the container but calling syscall.CLONE_NEWUTS
//NEWUTS is the Unix namespace for hostname
func (n *Namespace) GetHostName(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
}

//SetHostName for setting hostname
func (n *Namespace) SetHostName(hostname string, cmd *exec.Cmd) {
	syscall.Sethostname([]byte(hostname))
}

//SetDomainName for setting host domain name
func (n *Namespace) SetDomainName(cmd *exec.Cmd, hostname string) {
	syscall.Setdomainname([]byte(hostname))
}
