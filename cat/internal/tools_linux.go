package internal

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strings"
)

func (it *Environment) vendor() *Environment {
	fp := "/sys/class/dmi/id/sys_vendor"
	if FileExist(fp) {
		it.Vendor = strings.TrimSpace(String(&fp))
	}
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

func (it *Environment) release() *Environment {
	var NAME, VERSION, ID, ID_LIKE string
	var v = func(s string, l int) string {
		return strings.Trim(s[l:], "\"")
	}
	fp := "/etc/os-release"
	if !FileExist(fp) {
		return it
	}
	ls := strings.Split(String(&fp), "\n")
	for i := 0; i < len(ls); i++ {
		switch {
		case strings.HasPrefix(ls[i], "NAME="):
			NAME = v(ls[i], 5)
		case strings.HasPrefix(ls[i], "VERSION="):
			VERSION = v(ls[i], 8)
		case strings.HasPrefix(ls[i], "ID="):
			ID = v(ls[i], 3)
		case strings.HasPrefix(ls[i], "ID_LIKE="):
			ID_LIKE = v(ls[i], 8)
		}
	}
	it.Name = strings.Join([]string{NAME, VERSION}, " ")
	it.Platform = ID
	if _, has := releaseSet[ID]; !has {
		if ID_LIKE != "" {
			it.Platform = strings.Fields(ID_LIKE)[0]
		}
	}
	return it
}

func (it *Environment) android() *Environment {
	it.Platform = strings.ToLower(Commandline("", []string{"uname", "-o"}))
	if it.Platform == "android" {
		it.Processor = Commandline("", []string{"getprop", "ro.config.cpu_info_display"})
		it.Vendor = Commandline("", []string{"getprop", "ro.product.manufacturer"})
		it.Name = "Android " + Commandline("", []string{"getprop", "ro.system.build.version.release"})
	}
	return it
}

var releaseSet = map[string]byte{
	"fedora": 0,
	"rhel":   0,
	"centos": 0,
	"debian": 0,
	"ubuntu": 0,
}
