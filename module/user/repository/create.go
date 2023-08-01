package repository

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (r *UserRepository) Create(user *entity.User) error {
	err := r.db.Sql().DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
