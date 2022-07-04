package main

import (
	"fmt"
	"github.com/matsuwin/cat"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("demo> sysctl stop serviceName")
		return
	}
	if err := cat.Sysctl(os.Args[1], os.Args[2]); err != nil {
		cat.Stderr(err.Error())
	}
}
