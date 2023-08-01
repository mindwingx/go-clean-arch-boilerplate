package usecase

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (u UserUsecase) GetAll() (users []entity.User, err error) {
	users, err = u.userRepo.GetAll()
	if err != nil {
		//todo: implement call the log service method(sentry, zap-log, etc)
		return nil, err
	}

	return users, nil
}

func (u UserUsecase) GetById(id string) (user *entity.User, err error) {
	user, err = u.userRepo.GetById(id)
	if err != nil {
		//todo: implement call the log service method(sentry, zap-log, etc)
		return nil, err
	}

	return user, nil
}
