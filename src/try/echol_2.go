package try

import (
	"fmt"
	"os"
	"strings"
)

func echol2() {

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

}
