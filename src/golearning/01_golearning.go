package golearning

import (
	"bufio"
	"fmt"
	"os"
)

func golearning() {

	counts := make(map[string]int)

	filelist := os.Args[1:]

	for _, file := range filelist {
		input_fd, _ := os.Open(file)
		input_file := bufio.NewScanner(input_fd)
		for input_file.Scan() {
			counts[input_file.Text()]++
		}

		input_fd.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
