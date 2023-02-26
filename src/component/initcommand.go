package component

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
)

var MydockerInitCommand = cli.Command{
	Name:  "init",
	Usage: "use for mydocker init operations like mounting the /proc after create all subnamespace!",
	// Aliases: []string{"init"},
	Action: func(ctx *cli.Context) error {
		fmt.Printf("init command = %s \n", ctx.Args().Get(0))
		err := InitProcess()
		if err != nil {
			log.Fatal(err)
		}
		return err
	},
}

func InitProcess() error {
	fmt.Println("enter into InitProcess")

	cmdArray := readCommand()

	cmdPath, err := exec.LookPath(cmdArray[0])
	if err != nil {
		fmt.Printf("Exec loop path error %v", err)
		return err
	}

	// 开始配置mydocker内部的环境,目前只 mount /proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV

	// 需要先mount 根目录和 设置 private， 才能保证不同容器的/proc不是宿主机的/proc，不设置的话，宿主机需要mount /proc 后才能查看/proc
	if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		fmt.Printf("mount / fails: %v \n", err)
		return err
	}

	// 通过syscall.exec（）保证docker的进程号为1
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err := syscall.Exec(cmdPath, cmdArray[0:], os.Environ()); err != nil {
		fmt.Println("exec init command wrong !!")
		log.Fatal(err.Error())
	}

	return nil
}

func readCommand() []string {
	// 去掉stdin， stdout，stderr，默认index为3的就是创建的fd
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := io.ReadAll(pipe)
	if err != nil {
		fmt.Printf("init read pipe error %v", err)
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}
