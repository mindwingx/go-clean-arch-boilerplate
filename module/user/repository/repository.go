package repository

import (
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
	user "github.com/mindwingx/go-clean-arch-boilerplate/module/user/abstraction"
)

type UserRepository struct {
	db driver.SqlAbstraction
}

func NewUserRepo(db driver.SqlAbstraction) user.UserRepo {
	return &UserRepository{
		db: db,
	}
}
