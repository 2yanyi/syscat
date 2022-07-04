package main

import (
	"fmt"
	"r/pkg/cat"
)

func main() {
	info := cat.Syscat()
	fmt.Printf("%s\n", cat.Json(info))
}
