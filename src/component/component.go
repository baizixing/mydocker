package component

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

var MydockerRunCommand = cli.Command{
	Name:  "run",
	Usage: "use for run command, create all subnamespace",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "it",
			Usage: "enable false",
			Value: false,
		},
	},

	Action: func(ctx *cli.Context) error {
		// 检查输入是否有问题
		if ctx.NArg() < 1 {
			fmt.Printf("input args must more than 1.")
			return fmt.Errorf("input args must more than 1")
		}

		// 调用Run函数,开始command命令生成，
		command := ctx.Args().Get(0)

		fmt.Printf("it=%t   run command=%s \n", ctx.Bool("it"), command)

		Run(ctx.Bool("it"), command)

		return nil
	},
}

var MydockerInitCommand = cli.Command{
	Name:  "init",
	Usage: "use for mydocker init operations like mounting the /proc after create all subnamespace!",
	// Aliases: []string{"init"},
	Action: func(ctx *cli.Context) error {
		fmt.Printf("init command = %s \n", ctx.Args().Get(0))
		err := InitProcess(ctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		return err
	},
}

func InitProcess(command string) error {
	fmt.Println("enter into InitProcess")
	// 开始配置mydocker内部的环境,目前只 mount /proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV

	// 需要先mount 根目录和 设置 private， 才能保证不同容器的/proc不是宿主机的/proc，不设置的话，宿主机需要mount /proc 后才能查看/proc
	if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		fmt.Printf("mount / fails: %v \n", err)
		return err
	}

	fmt.Println("...............sleep begin.................")
	time.Sleep(time.Duration(10) * time.Second)

	// 通过syscall.exec（）保证docker的进程号为1
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err := syscall.Exec("greet", []string{command}, os.Environ()); err != nil {
		fmt.Println("exec init command wrong !!")
		log.Fatal(err.Error())
	}

	return nil
}

func CreateNewProcess(itFlag bool, command string) *exec.Cmd {
	// 通过exec.Command方式调用init
	params := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", params...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,

		// UidMappings和GidMappings用来配置user namespace
		// UidMappings: []syscall.SysProcIDMap{
		// 	{ContainerID: 0, HostID: 0, Size: 1},
		// },

		// GidMappings: []syscall.SysProcIDMap{
		// 	{ContainerID: 0, HostID: 0, Size: 1},
		// },
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

/*
run command, 步骤如下：
 1. 宿主机fork一个新的进程，之后挂载 /proc目录，并配置
 3. 判断it flag是否为true，true则将容器进程后台运行
*/
func Run(itFlag bool, command string) {

	fmt.Println("enter into Run function")

	cmd := CreateNewProcess(itFlag, command)

	fmt.Println("cmd start .....")
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
	os.Exit(-1)
}
