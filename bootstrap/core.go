package bootstrap

import (
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
)

func (a *App) initRegistry() {
	a.registry = driver.NewViper(a.locale)
	a.registry.InitRegistry()
}

func (a *App) initLocale() {
	a.locale = driver.NewLocale(a.registry.ValueOf("locale"))
	a.locale.InitLocale()
}

func (a *App) initDatabase() {
	a.database = driver.NewSql(a.registry.ValueOf("sql"), a.locale)
	a.database.InitSql()
	a.database.Migrate()
}

func (a *App) initCache() {
	a.cache = driver.NewRedis(a.registry.ValueOf("cache"), a.locale)
	a.cache.InitCache()
}

func (a *App) initServices() {
	// rpc
	a.rpc = driver.NewRpc(a.registry.ValueOf("rpc"), a.locale)
	a.rpc.InitRpcService()

	// api
	a.http = driver.NewGin(a.registry.ValueOf("http"), a.locale)
	a.http.InitHttp()

	// todo : init the gRPC, Cron Job, 3rd-party services, etc
}
