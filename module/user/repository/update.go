package repository

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (r *UserRepository) Update(user *entity.User, id string) (err error) {
	err = r.db.Sql().DB.Model(user).Update("id", id).Error

	if err != nil {
		return err
	}

	return nil
}
