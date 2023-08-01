package repository

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (r *UserRepository) GetAll() ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetById(id string) (user *entity.User, err error) {
	err = r.db.Sql().DB.Where("id  = ?", id).First(&user).Error
	if err != nil {
		//todo: handle logger
		return nil, err
	}

	return user, nil
}
