package subsystem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"utils"
)

type memorySubsystem struct {
}

func (mem memorySubsystem) Name() string {
	return "memory"
}

func (mem memorySubsystem) Setlimit(cgroupName string, resconfig *SubsysResconfig) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(mem.Name(), cgroupName, true); err == nil {
		if resconfig.Memory != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(resconfig.Memory), 0644); err != nil {
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	} else {
		return nil
	}
}

func (mem memorySubsystem) AddTaskToCgroups(cgroupName string, pid int) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(mem.Name(), cgroupName, false); err == nil {
		if err := os.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupName, err)
	}
}

func (mem memorySubsystem) Remove(cgroupName string) error {
	if subsysCgroupPath, err := utils.GetCgroupPath(mem.Name(), cgroupName, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}
