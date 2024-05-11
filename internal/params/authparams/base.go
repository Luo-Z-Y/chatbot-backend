package authparams

import "backend/internal/model"

type Params struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func (p *Params) ToModel() *model.User {
	return &model.User{Username: p.Username, EncryptedPassword: p.Password}
}
