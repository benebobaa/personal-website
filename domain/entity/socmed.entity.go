package entity

import "personal-website/delivery/http/dto/response"

type Socmed struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;not null"`
	Link string `gorm:"column:link;not null"`
}

func (Socmed) TableName() string {
	return "socmeds"
}

func ToSocmedResponses(socmeds []Socmed) []response.SocmedResponse {
	var socmedResponses []response.SocmedResponse
	for _, socmed := range socmeds {
		socmedResponses = append(socmedResponses, response.SocmedResponse{
			ID:   socmed.ID,
			Name: socmed.Name,
			Link: socmed.Link,
		})
	}
	return socmedResponses
}
