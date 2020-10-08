package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/echenim/microcloud/utils"

	"github.com/echenim/microcloud/core/namespaces"
)

//Parent function for accessing namespaces and cgroups
//go run main.go run /bin/bash
func Parent() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())
	nspace := namespaces.Namespace{}
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	nspace.GetInput(cmd)
	nspace.GetOutput(cmd)
	nspace.GetError(cmd)
	nspace.GetHostName(cmd)
	nspace.GetProcessID(cmd)
	r := cmd.Run()
	utils.Check(r, "FAILED")
}
