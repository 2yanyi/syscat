package internal

import (
	"errors"
	"fmt"
	"github.com/kardianos/service"
)

type Program struct{}

func (r *Program) Start(s service.Service) error { return nil }

func (r *Program) Stop(s service.Service) error { return nil }

func Sysctl(serviceAction, serviceName string) error {
	app := &Program{}
	handle, err := service.New(app, &service.Config{Name: serviceName})
	if err != nil {
		return err
	}
	switch serviceAction {
	case "start":
		err = handle.Start()
	case "stop":
		err = handle.Stop()
	case "restart":
		err = handle.Restart()
	case "uninstall":
		if err = handle.Stop(); err != nil {
			return err
		} else {
			err = handle.Uninstall()
		}
	case "status":
		code, _ := handle.Status()
		return errors.New(fmt.Sprintf("%d", code))
	}
	return err
}
