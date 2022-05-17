package internal

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strings"
)

func (it *Environment) vendor() *Environment {
	it.Vendor = strings.TrimSpace(String("/sys/class/dmi/id/sys_vendor"))
	return it
}

func (it *Environment) kernel() *Environment {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		it.Kernel = "unknown"
		return it
	}
	it.Kernel = fmt.Sprintf("%s", uname.Release)
	return it
}

func (it *Environment) release() {
	var NAME, VERSION, ID, ID_LIKE string
	var v = func(s string, l int) string {
		return strings.Trim(s[l:], "\"")
	}
	for _, elem := range strings.Split(String("/etc/os-release"), "\n") {
		switch {
		case strings.HasPrefix(elem, "NAME="):
			NAME = v(elem, 5)
		case strings.HasPrefix(elem, "VERSION="):
			VERSION = v(elem, 8)
		case strings.HasPrefix(elem, "ID="):
			ID = v(elem, 3)
		case strings.HasPrefix(elem, "ID_LIKE="):
			ID_LIKE = v(elem, 8)
		}
	}
	it.Name = strings.Join([]string{NAME, VERSION}, " ")
	it.Platform = ID
	if _, has := releaseSet[ID]; !has {
		if ID_LIKE != "" {
			it.Platform = strings.Fields(ID_LIKE)[0]
		}
	}

}

var releaseSet = map[string]*struct{}{
	"fedora": nil,
	"rhel":   nil,
	"centos": nil,
	"debian": nil,
	"ubuntu": nil,
}
