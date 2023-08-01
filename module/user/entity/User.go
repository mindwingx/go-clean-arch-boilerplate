package entity

import (
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	NickName  string    `json:"nick_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) Table() string {
	return helper.UsersTable
}
