package bootstrap

import "github.com/mindwingx/go-clean-arch-boilerplate/module/user"

func (a *App) initUserModule() {
	a.UserModule = user.NewModule(a.database, a.locale)
}
