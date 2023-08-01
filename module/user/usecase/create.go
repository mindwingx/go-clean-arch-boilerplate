package usecase

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (u UserUsecase) Create(user *entity.User) (err error) {
	err = u.userRepo.Create(user)
	if err != nil {
		//todo: implement call the log service method(sentry, zap-log, etc)
		return err
	}

	return nil
}
