package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"utils"
)

type cpusetSubsystem struct {
}

func (cpuset cpusetSubsystem) Name() string {
	return "cpuset"
}

func (cpuset cpusetSubsystem) Setlimit(cgroupName string, resconfig *SubsysResconfig) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpuset.Name(), cgroupName, true); err == nil {
		if resconfig.Cpuset != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(resconfig.Cpuset), 0644); err != nil {
				return fmt.Errorf("set cgroup cpuset fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (cpuset cpusetSubsystem) AddTaskToCgroups(cgroupName string, pid int) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpuset.Name(), cgroupName, false); err == nil {
		if err := os.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupName, err)
	}
}

func (cpuset cpusetSubsystem) Remove(cgroupName string) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpuset.Name(), cgroupName, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}
