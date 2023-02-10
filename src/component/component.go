package component

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli/v2"
)

var MydockerRunCommand = cli.Command{
	Name:    "runCommand",
	Usage:   "use for run command, create all subnamespace",
	Aliases: []string{"run"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "it",
			Value: false,
		},
	},

	Action: func(ctx *cli.Context) error {
		// 开始配置namespace
		// 检查输入是否有问题
		if ctx.NArg() < 1 {
			fmt.Printf("input args must more than 1.")
			return fmt.Errorf("input args must more than 1")
		}

		// 调用Run函数
		command := ctx.Args().Get(0)
		Run(ctx.Bool("it"), command)
		return nil
	},
}

var MydockerInitCommand = cli.Command{
	Name:    "runCommand",
	Usage:   "use for mydocker init operations like mounting the /proc after create all subnamespace!",
	Aliases: []string{"init"},
	Action: func(ctx *cli.Context) error {

		return nil
	},
}

/*
run command, 步骤如下：
 1. 宿主机fork一个新的进程，之后挂载 /proc目录，并配置
 3. 判断it flag是否为true，true则将容器进程后台运行
*/
func Run(itFlag bool, command string) *exec.Cmd {

	// 调用
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,

		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: 0, Size: 1},
		},

		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: 0, Size: 1},
		},
	}

	if itFlag {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Println("it Flag is true")
	} else {
		fmt.Println("it flag is false")
	}
	return cmd
}
