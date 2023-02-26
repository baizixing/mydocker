package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"utils"
)

type cpuSubsystem struct {
}

func (cpushare cpuSubsystem) Name() string {
	return "cpu"
}

func (cpushare cpuSubsystem) Setlimit(cgroupName string, resconfig *SubsysResconfig) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpushare.Name(), cgroupName, true); err == nil {
		if resconfig.Cpushare != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "cpu.shares"), []byte(resconfig.Cpushare), 0644); err != nil {
				return fmt.Errorf("set cgroup cpu share fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (cpushare cpuSubsystem) AddTaskToCgroups(cgroupName string, pid int) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpushare.Name(), cgroupName, false); err == nil {
		if err := os.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupName, err)
	}
}

func (cpushare cpuSubsystem) Remove(cgroupName string) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(cpushare.Name(), cgroupName, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}
