package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"github.com/mindwingx/go-clean-arch-boilerplate/module/user/abstraction"
	"net/http"
)

type FetchUser struct {
	userUc abstraction.UserUc
	locale driver.LocaleAbstraction
}

func NewFetchUser(userUc abstraction.UserUc, locale driver.LocaleAbstraction) gin.HandlerFunc {
	handler := &FetchUser{
		userUc: userUc,
		locale: locale,
	}
	return handler.FetchUser
}

func (fu FetchUser) FetchUser(c *gin.Context) {
	//todo: use validation here

	id := c.Param("id")
	user, err := fu.userUc.GetById(id)

	if err != nil {
		helper.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, map[string]interface{}{"user": user})
	return
}
