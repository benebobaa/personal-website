package entity

import "personal-website/delivery/http/dto/response"

type Admin struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Username string `gorm:"column:username;not null;unique"`
	Password string `gorm:"column:password;not null"`
}

func (Admin) TableName() string {
	return "admins"
}

func (a *Admin) ToResponse() *response.AdminResponse {
	return &response.AdminResponse{
		ID:       a.ID,
		Username: a.Username,
	}
}
