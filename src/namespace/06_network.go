package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func network() {
	cmd := exec.Command("sh") //创建sh
	//配置sh初始参数，开启UTS,IPC,PID,Mount,User,Network
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,

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
