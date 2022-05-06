package main

import (
	"fmt"
	"r/cat"
)

func main() {
	info := cat.SystemInfo()
	fmt.Printf("%s\n", cat.Json(info))
}
