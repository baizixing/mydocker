package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func open_uts() {
	cmd := exec.Command("sh")               //创建sh
	cmd.SysProcAttr = &syscall.SysProcAttr{ //配置sh初始参数，开启uts
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
