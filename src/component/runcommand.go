package component

import (
	"cgroups"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"subsystem"
	"syscall"

	"github.com/urfave/cli/v2"
)

var MydockerRunCommand = cli.Command{
	Name:  "run",
	Usage: "use for run command, create all subnamespace",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "memory",
			Aliases: []string{"m"},
			Usage:   "memory limit",
		},
		&cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		&cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
	},

	Action: func(ctx *cli.Context) error {
		// 检查输入是否有问题
		if ctx.NArg() < 1 {
			fmt.Printf("input args must more than 1.")
			return fmt.Errorf("input args must more than 1")
		}

		var cmdArrayForInit []string
		// 调用Run函数,开始command命令生成，
		for index := 0; index < ctx.Args().Len(); index++ {
			cmdArrayForInit = append(cmdArrayForInit, ctx.Args().Get(index))
			fmt.Println(cmdArrayForInit)
		}

		tty := ctx.Bool("it")

		resconf := &subsystem.SubsysResconfig{
			Memory:   ctx.String("m"),
			Cpuset:   ctx.String("cpuset"),
			Cpushare: ctx.String("cpushare"),
		}

		run(tty, cmdArrayForInit, resconf)

		return nil
	},
}

/*
run command, 步骤如下：
 1. 配置一个新的终端进程，包括namespace的系统参数设置，是否需要隐藏新终端进程的输出等
 2. 终端启动后，开始初始化该进程内的环境，包括/proc挂载等，之后再执行command
*/
func run(itFlag bool, commandArray []string, resconf *subsystem.SubsysResconfig) {

	fmt.Println("enter into Run function")

	cmd, writePipe := CreateNewProcess(itFlag, commandArray)

	fmt.Println("cmd start .....")
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup")
	defer cgroupManager.Remove()

	cgroupManager.Setlimit(resconf)
	cgroupManager.AddTaskToCgroups(cmd.Process.Pid)
	sendInitCommand(commandArray, writePipe)

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
	os.Exit(-1)
}

func CreateNewProcess(itFlag bool, command []string) (*exec.Cmd, *os.File) {

	readPipe, writePipe, err := NewPipe()
	if err != nil {
		fmt.Printf("New pipe error %v", err)
		return nil, nil
	}
	// 通过exec.Command方式调用init
	cmd := exec.Command("/proc/self/exe", "init")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}

	if itFlag {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Println("it Flag is true")
	} else {
		fmt.Println("it flag is false")
	}

	cmd.ExtraFiles = []*os.File{readPipe}
	return cmd, writePipe
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}

	return read, write, nil
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	fmt.Printf("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
