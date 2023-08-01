package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user/abstraction"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user/delivery"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user/repository"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user/usecase"
)

type Module struct {
	UserRepo abstraction.UserRepo
	UserUc   abstraction.UserUc
	Fetch    gin.HandlerFunc
}

func NewModule(db driver.SqlAbstraction, locale driver.LocaleAbstraction) *Module {
	m := new(Module)
	m.UserRepo = repository.NewUserRepo(db)
	m.UserUc = usecase.NewUserUsecase(m.UserRepo)
	m.Fetch = delivery.NewFetchUser(m.UserUc, locale)
	return m
}
