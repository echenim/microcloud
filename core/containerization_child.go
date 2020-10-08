package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/echenim/microcloud/core/namespaces"
	"github.com/echenim/microcloud/utils"
)

//Child function that help fork exec
func Child() {

	nspace := namespaces.Namespace{}
	//hostname, _ := os.Hostname()
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	nspace.GetInput(cmd)
	nspace.GetOutput(cmd)
	nspace.GetError(cmd)
	nspace.SetHostName("Quants", cmd)
	//dir := syscall.Chroot("/home/itachi/ubuntufs")
	//rt := syscall.Chdir("/")
	//utils.Check(dir, "FAILED")
	//utils.Check(rt, "FAILED")
	// utils.Check(syscall.Chroot("/home/itachi/rootfs"), "FAILED")
	// utils.Check(os.Chdir("/"), "FAILED")
	// utils.Check(syscall.Mount("proc", "proc", "proc", 0, ""), "FAILED")

	r := cmd.Run()
	utils.Check(r, "FAILED")
}
