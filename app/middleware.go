package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mindwingx/go-clean-arch-boilerplate/app/middleware"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"net/http"
)

func (a *App) CheckAuth(c *gin.Context) {
	if err := middleware.CheckAuth(c, a.locale); err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Next()
}
