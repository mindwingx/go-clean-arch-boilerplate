package bootstrap

import (
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user"
)

type App struct {
	registry driver.RegistryAbstraction
	database driver.SqlAbstraction
	cache    driver.CacheAbstraction
	http     driver.HttpAbstraction
	rpc      driver.RpcAbstraction
	locale   driver.LocaleAbstraction
	// Service Modules
	UserModule *user.Module
}

func New() *App {
	return &App{}
}

func (a *App) InitDrivers() {
	a.initRegistry()
	a.initLocale()
	a.initDatabase()
	a.initCache()
	a.initServices()
}

func (a *App) InitModules() {
	a.initUserModule()
}
