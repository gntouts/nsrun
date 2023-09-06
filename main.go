package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/vishvananda/netns"
)

func main() {
	args := os.Args
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			usage()
			os.Exit(0)
		}
	}

	if len(args) == 1 {
		usage()
		os.Exit(0)
	}

	if len(args) < 3 {
		fmt.Println("ERR: too few arguments")
		usage()
		os.Exit(1)
	}

	targetPid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("ERR: failed to convert given PID to int with error: %v\n", err)
		os.Exit(1)
	}

	pidExists, err := process.PidExists(int32(targetPid))
	if err != nil {
		fmt.Printf("ERR: failed to check if given PID exists with error: %v\n", err)
		os.Exit(1)
	}
	if !pidExists {
		fmt.Println("ERR: given PID does not exist")
		os.Exit(1)
	}

	targetNetNS, err := netns.GetFromPid(int(targetPid))
	if err != nil {
		fmt.Printf("ERR: failed to find network namespace for given PID with error: %v\n", err)
		os.Exit(1)
	}

	err = netns.Set(targetNetNS)
	if err != nil {
		fmt.Printf("ERR: failed to enter the network namespace for given PID with error: %v\n", err)
		os.Exit(1)
	}

	binary := args[2]
	binaryPath, err := exec.LookPath(binary)
	if err != nil {
		fmt.Printf("ERR: failed to find the given binary \"%s\" with error: %v\n", binary, err)
		os.Exit(1)
	}
	cmdArgs := args[2:]
	env := os.Environ()
	err = syscall.Exec(binaryPath, cmdArgs, env)
	if err != nil {
		fmt.Printf("ERR: failed to EXECVE the given command with error: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	usageMsg := `NAME:
	execns - execute command in given process's network namespace
			 
 
 USAGE:
 	execns [PID] [COMMAND]`
	fmt.Println(usageMsg)
}
