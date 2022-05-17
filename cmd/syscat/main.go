package main

import (
	"fmt"
	"github.com/matsuwin/cat"
)

func main() {
	info := cat.SystemInfo()
	fmt.Println(info.SpeedIconTitle())
	fmt.Printf("%s\n", cat.Json(info))
}
