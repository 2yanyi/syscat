package internal

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strings"
)

func vendor() string {
	return strings.TrimSpace(String("/sys/class/dmi/id/sys_vendor"))
}

func kernel() string {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "unknown"
	}
	return fmt.Sprintf("%s", uname.Release)
}

func release() (title, platform string) {
	var name, version string
	for _, elem := range strings.Split(String(releaseFp), "\n") {
		switch {
		case strings.HasPrefix(elem, releaseIDLike):
			platform = strings.Fields(strings.Trim(elem[len(releaseIDLike):], _tag))[0]
		case strings.HasPrefix(elem, releaseVersion):
			version = strings.Trim(elem[len(releaseVersion):], _tag)
		case strings.HasPrefix(elem, releaseName):
			name = strings.Trim(elem[len(releaseName):], _tag)
		}
	}
	title = strings.Join([]string{name, version}, " ")
	return
}

const releaseFp = "/etc/os-release"
const releaseIDLike = "ID_LIKE="
const releaseVersion = "VERSION="
const releaseName = "NAME="
const _tag = "\""
