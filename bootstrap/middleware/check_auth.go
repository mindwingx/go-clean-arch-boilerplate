package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mindwingx/go-clean-arch-boilerplate/driver"
)

func CheckAuth(c *gin.Context, l driver.LocaleAbstraction) (err error) {
	_ = c.GetHeader("jwt-token")

	//todo: validate the JWT token here

	if err != nil {
		err = errors.New(l.Get("unauthorized"))
		return err
	}

	return nil
}
