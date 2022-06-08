package internal

import (
	"github.com/jaypipes/ghw"
	"strings"
)

func graphics() []string {
	drivers := make([]string, 0)
	info, err := ghw.GPU()
	if err != nil {
		Stderr(err.Error())
	}
	if info == nil {
		return nil
	}
	for i := 0; i < len(info.GraphicsCards); i++ {
		if info.GraphicsCards[i] == nil {
			continue
		}
		if info.GraphicsCards[i].DeviceInfo == nil {
			continue
		}
		if info.GraphicsCards[i].DeviceInfo.Product == nil {
			continue
		}
		if info.GraphicsCards[i].DeviceInfo.Product.Name == "SVGA II Adapter" {
			continue
		}
		if strings.Contains(info.GraphicsCards[i].DeviceInfo.Product.Name, "Graphics") {
			continue
		}
		drivers = append(drivers, info.GraphicsCards[i].DeviceInfo.Product.Name)
	}
	return drivers
}
