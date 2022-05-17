package internal

import (
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
)

func (it *Environment) vendor() *Environment {
	if product, _ := ghw.Product(); product != nil {
		it.Vendor = product.Vendor
	}
	return it
}

func (it *Environment) kernel() *Environment {
	return it
}

func (it *Environment) release() {
	info, err := host.Info()
	if err != nil {
		it.Name = "unknown"
		it.Kernel = "unknown"
		return
	}
	it.Name = info.Platform
	it.Kernel = info.KernelVersion
}
