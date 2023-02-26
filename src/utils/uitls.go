package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func FindHierarchyMountPoint(subsysName string) string {
	// 用于获取Hierarchy path, mountinfo一行最后几个字符是挂载类型
	// 比如 ,cpuset, 一行中按空格切分，第五个(index为4)的是实际挂载路径
	// 这里只有hierarchy的路径，还需要在下面创建Cgroup的文件夹

	f, err := os.Open("/proc/self/mountinfo")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	if _, err := f.Stat(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		temp_line := strings.Split(txt, " ")
		for _, opt := range strings.Split(temp_line[len(temp_line)-1], ",") {
			if opt == subsysName {
				return temp_line[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""

}

func GetCgroupPath(subsysName string, cgroupName string, autoCreate bool) (string, error) {
	hierarchyPath := FindHierarchyMountPoint(subsysName)
	_, err := os.Stat(path.Join(hierarchyPath, cgroupName))
	if (err == nil) || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(hierarchyPath, cgroupName), 0755); err == nil {

			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(hierarchyPath, cgroupName), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}
