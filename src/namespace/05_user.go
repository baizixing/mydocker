package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func user() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER, //开启UTS,IPC,PID,Mount,User

		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: 0, Size: 1},
		},

		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: 0, Size: 1},
		},
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
