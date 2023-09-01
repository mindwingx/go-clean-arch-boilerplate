package main

import "github.com/mindwingx/go-clean-arch-boilerplate/bootstrap"

func main() {
	service := bootstrap.New()
	service.InitDrivers()
	service.InitModules()
	service.Start()
}
