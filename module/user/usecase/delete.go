package usecase

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user/entity"

func (u UserUsecase) Delete(user *entity.User, id string) (err error) {
	err = u.userRepo.Delete(user, id)
	if err != nil {
		//todo: implement call the log service method(sentry, zap-log, etc)
		return err
	}

	return nil
}
