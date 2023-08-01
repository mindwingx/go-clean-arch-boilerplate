package driver

import (
	"fmt"
	src "github.com/mindwingx/go-clean-arch-boilerplate"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	registry "github.com/spf13/viper"
	"os"
)

type RegistryAbstraction interface {
	InitRegistry()
	ValueOf(string) RegistryAbstraction
	Parse(interface{})
}

type Viper struct {
	locale   LocaleAbstraction
	registry *registry.Viper
}

func NewViper(locale LocaleAbstraction) RegistryAbstraction {
	return &Viper{
		locale:   locale,
		registry: registry.New(),
	}
}

func (v *Viper) InitRegistry() {
	v.registry.SetConfigType("yml")

	configFile := fmt.Sprintf("%s/config/env.yml", src.Root())
	configBytes, err := os.Open(configFile)

	if err != nil {
		helper.CustomPanic("failed to load registry file", err)
	}

	err = v.registry.ReadConfig(configBytes)
	if err != nil {
		helper.CustomPanic("failed to read registry config file", err)
	}
}

func (v *Viper) ValueOf(key string) RegistryAbstraction {
	return &Viper{registry: v.registry.Sub(key)}
}

func (v *Viper) Parse(item interface{}) {
	err := v.registry.Unmarshal(&item)
	if err != nil {
		helper.CustomPanic("failed to retrieve registry configs", err)
	}
}
