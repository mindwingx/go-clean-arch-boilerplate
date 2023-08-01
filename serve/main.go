package main

import "github.com/mindwingx/go-clean-arch-boilerplate/app"

func main() {
	service := app.New()
	service.InitDrivers()
	service.InitModules()
	service.Start()
}
