package subsystem

type SubsysResconfig struct {
	// 这里记录的只是用户命令行输入的值，作为参数传给setlimit()函数，
	Memory   string
	Cpuset   string
	Cpushare string
}

// 因为对memory, cpuset, cpu三个subsystem，需要进行的操作是一样的，所以可以通过接口来实现
type Subsystem interface {
	Name() string
	Setlimit(path string, resconfig *SubsysResconfig) error
	AddTaskToCgroups(path string, pid int) error
	Remove(path string) error
}

// 创建三个subsystem的实例，分别用于处理memory，cpuset，cpu
var SubsysIns = []Subsystem{
	&memorySubsystem{},
	&cpusetSubsystem{},
	&cpuSubsystem{},
}
