package driver

import (
	"fmt"
	"github.com/fatih/color"
	core "github.com/gin-gonic/gin"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"time"
)

type HttpAbstraction interface {
	InitHttp()
	StartHttp()
	Service() *Gin
}

type (
	Gin struct {
		config apiConfig
		locale LocaleAbstraction
		Core   *core.Engine
	}

	apiConfig struct {
		Host            string
		Port            string
		Development     bool
		ShutdownTimeout time.Duration
	}
)

func NewGin(registry RegistryAbstraction, locale LocaleAbstraction) HttpAbstraction {
	engin := new(Gin)
	registry.Parse(&engin.config)
	engin.locale = locale
	engin.Core = core.New()
	return engin
}

func (g *Gin) InitHttp() {
	err := g.Core.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		helper.CustomPanic(g.locale.Get("http_trusted_proxy"), err)
	}

	if g.config.Development == false {
		core.SetMode(core.ReleaseMode)
	}

	g.Core.Use(core.Logger())
	g.Core.Use(core.Recovery())
}

func (g *Gin) StartHttp() {
	address := fmt.Sprintf("%s:%s", g.config.Host, g.config.Port)
	color.Cyan(g.locale.Get("http_start"))

	err := g.Core.Run(address)
	if err != nil {
		helper.CustomPanic(g.locale.Get("http_start_failure"), err)
	}
}

func (g *Gin) Service() *Gin {
	return g
}
