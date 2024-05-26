package viewmodel

import (
	"backend/internal/model"
)

type UserView struct {
	ID       uint
	Username string
}

func UserViewFrom(u *model.User) *UserView {
	return &UserView{
		ID: u.ID, Username: u.Username,
	}
}
