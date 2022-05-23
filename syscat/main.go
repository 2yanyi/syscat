package main

import (
	"fmt"
	"github.com/matsuwin/cat"
)

func main() {
	info := cat.SystemInfo()
	fmt.Printf("%s\n", cat.Json(info))
}
