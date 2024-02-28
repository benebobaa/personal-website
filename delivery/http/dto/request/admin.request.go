package request

import "personal-website/domain/entity"

type AdminRequest struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (a *AdminRequest) ToEntity() *entity.Admin {
	return &entity.Admin{
		Username: a.Username,
		Password: a.Password,
	}
}
