package repository

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (r *UserRepository) Delete(user *entity.User, id string) error {
	err := r.db.Sql().DB.Delete(user, id).Error
	if err != nil {
		return err
	}
	return nil
}
