// Building a container in Go - a talk given by Liz Rice
// https://youtu.be/Utf-A4rODH8

package main

import (
	"fmt"       // format
	"os"        // a platform independent library for OS functionality. https://golang.org/pkg/os/
	"os/exec"   //  runs external commands - wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.  https://golang.org/pkg/os/exec/
	"syscall"
)

// docker run < container-name > { ... cmd args ... }
// go run < main.go > { ... run cmd args ... }
func main() {
	switch os.Args[1]{
	case "run":
		run()
	default:
		panic("what??")
	}
}

func run () {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS,          //UTS - Unix Time Sharing System
	}

	must (cmd.Run())
}

func must (err error) {
	if err != nil {
		panic(err)
	}
}

