package entity

import "personal-website/delivery/http/dto/response"

type Hero struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName   string `gorm:"column:first_name;not null"`
	LastName    string `gorm:"column:last_name;not null"`
	Description string `gorm:"column:description;not null"`
}

func (Hero) TableName() string {
	return "heroes"
}

func (h *Hero) ToResponse() *response.HeroResponse {
	return &response.HeroResponse{
		ID:          h.ID,
		FirstName:   h.FirstName,
		LastName:    h.LastName,
		Description: h.Description,
	}
}
