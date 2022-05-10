package internal

import (
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
)

func vendor() (_ string) {
	if product, _ := ghw.Product(); product != nil {
		return product.Vendor
	}
	return
}

func kernel() string {
	return ""
}

func release() (title, platform string) {
	info, err := host.Info()
	if err != nil {
		return "unknown", "unknown"
	}
	return info.Platform, info.KernelVersion
}
