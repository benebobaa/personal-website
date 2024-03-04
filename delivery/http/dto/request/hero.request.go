package request

import "personal-website/domain/entity"

type HeroRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type HeroUpdateRequest struct {
	ID          int    `json:"id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *HeroRequest) ToEntity() *entity.Hero {
	return &entity.Hero{
		FirstName:   h.FirstName,
		LastName:    h.LastName,
		Description: h.Description,
	}
}

func (h *HeroUpdateRequest) ToEntity() *entity.Hero {
	return &entity.Hero{
		ID:          h.ID,
		FirstName:   h.FirstName,
		LastName:    h.LastName,
		Description: h.Description,
	}
}
