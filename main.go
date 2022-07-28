package main

import (
	"fmt"
	"github.com/matsuwin/syscat/cat"
	"os"
)

func main() {
	args := len(os.Args)
	if args == 1 {
		fmt.Printf("%s\n", cat.Json(cat.Syscat()))
		return
	}
	if args < 3 {
		fmt.Println("  syscat stop serviceName")
		return
	}
	if err := cat.Sysctl(os.Args[1], os.Args[2]); err != nil {
		cat.Stderr(err.Error())
	}
}
