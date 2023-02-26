package cgroups

import (
	"subsystem"
)

type CgroupsManager struct {
	path string // 注意，因为要实现的是类似docker run name这样的命令，所以从命令行能拿到的只有当前的项目名。
	// 为了实现cgroups，我们需要找到memory，cpuset，cpushare三个hierarchy的位置，
	// 也就是这三个挂载点路径，然后在这个目录下mkdir name来创建子cgroup，最后加入limit和taskid实现
	resourceConfig subsystem.SubsysResconfig
}

func NewCgroupManager(path string) *CgroupsManager {
	return &CgroupsManager{
		path: path,
	}
}

func (cgroup CgroupsManager) Setlimit(resconf *subsystem.SubsysResconfig) error {
	for _, subsys := range subsystem.SubsysIns {
		subsys.Setlimit(cgroup.path, resconf)
	}
	return nil
}

func (cgroup CgroupsManager) AddTaskToCgroups(pid int) error {
	// 给上层提供方法，这里的path是项目名，不是真正的cgroup所在的路径
	for _, subsys := range subsystem.SubsysIns {
		subsys.AddTaskToCgroups(cgroup.path, pid)
	}
	return nil
}

func (cgroup CgroupsManager) Remove() error {
	// 进程结束前需要清理掉 hierarchy下的cgroup信息，目前采用的是简单的删除hierarchy下的所有cgroup文件夹
	for _, subsys := range subsystem.SubsysIns {
		subsys.Remove(cgroup.path)
	}
	return nil
}
