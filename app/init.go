package app

import (
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	a.registerRoutes()

	go func() {
		a.rpc.StartRpc()
	}()

	go func() {
		a.http.StartHttp()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c

	color.Yellow(a.locale.Get("service_shutdown"))
}
