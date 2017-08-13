package main

import (
	"os"

	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 代码写在这儿
}

func (p *program) Stop(s service.Service) error {
	return nil
}

/**
 * MAIN函数，程序入口
 */
func main() {
	svcConfig := &service.Config{
		Name:        "golangServiceDemo",     //服务显示名称
		DisplayName: "golangServiceDemoDisp", //服务名称
		Description: "Go Lang sercice",       //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			logger.Println("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			logger.Println("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
