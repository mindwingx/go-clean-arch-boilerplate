package app

func (a *App) registerRoutes() {
	http := a.http.Service().Core

	api := http.Group("api")

	user := api.Group("user")
	user.GET("/all", a.CheckAuth, a.UserModule.Fetch)
	user.GET("/:id", a.CheckAuth, a.UserModule.Fetch)
}
