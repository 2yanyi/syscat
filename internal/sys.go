package internal

import (
	"os/exec"
	"runtime"
	"strings"
)

type Environment struct {
	Vendor     string   `json:"vendor"`
	Name       string   `json:"name"`
	Perf       string   `json:"perf"`
	Processor  string   `json:"processor"`
	Graphics   string   `json:"graphics,omitempty"`
	Platform   string   `json:"platform"`
	Kernel     string   `json:"kernel"`
	Init       string   `json:"init,omitempty"`
	LanAddress []string `json:"lanAddress,omitempty"`
}

func SystemInfo() *Environment {
	it := &Environment{}
	it.Perf, it.Processor = cpuTitle()
	it.Graphics = strings.Join(graphics(), ", ")
	it.Vendor = vendor()

	switch runtime.GOOS {
	case "windows":
		{
			it.Platform = runtime.GOOS
			it.Name, it.Kernel = release()
			it.Kernel = "NT " + strings.Fields(it.Kernel)[0]
		}
	case "linux":
		{
			it.Name, it.Platform = release()
			it.Kernel = kernel()
			it.Kernel = "Linux " + strings.Split(it.Kernel, "-")[0]
			if fp, _ := exec.LookPath("systemctl"); fp != "" {
				it.Init = "systemd"
			} else if fp, _ = exec.LookPath("service"); fp != "" {
				it.Init = "upstart" // sysvinit
			} else {
				it.Init = "no init"
			}
		}
	}
	it.LanAddress = LanAddress()

	return it
}
