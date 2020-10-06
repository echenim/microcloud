package core

import (
	"os"
	"os/exec"

	"github.com/echenim/microcloud/core/namespaces"
)

//Parent function for accessing namespaces and cgroups
//go run main.go run /bin/bash
func Parent() {

	nspace := namespaces.Namespace{}
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	nspace.GetHostName(cmd)

}
