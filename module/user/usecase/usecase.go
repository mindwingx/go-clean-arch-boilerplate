package usecase

import (
	user "github.com/mindwingx/go-clean-arch-boilerplate/module/user/abstraction"
)

type UserUsecase struct {
	userRepo user.UserRepo
}

func NewUserUsecase(repo user.UserRepo) user.UserUc {
	return &UserUsecase{userRepo: repo}
}
