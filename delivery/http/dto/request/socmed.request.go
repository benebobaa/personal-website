package request

import "personal-website/domain/entity"

type SocmedRequest struct {
	Name string `json:"name" validate:"required"`
	Link string `json:"link" validate:"required"`
}

type SocmedUpdateRequest struct {
	ID   int    `json:"id" validate:"required"`
	Link string `json:"link" validate:"required"`
}

func (s *SocmedRequest) ToEntity() *entity.Socmed {
	return &entity.Socmed{
		Name: s.Name,
		Link: s.Link,
	}
}

func (s *SocmedUpdateRequest) ToEntity() *entity.Socmed {
	return &entity.Socmed{
		ID:   s.ID,
		Link: s.Link,
	}
}
