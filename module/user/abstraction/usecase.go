package abstraction

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

type UserUc interface {
	Create(user *entity.User) error
	Update(user *entity.User, id string) error
	Delete(user *entity.User, id string) error
	GetAll() ([]entity.User, error)
	GetById(id string) (*entity.User, error)
}
