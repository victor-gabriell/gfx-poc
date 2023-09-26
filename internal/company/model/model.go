package model

import "github.com/google/uuid"

type Company struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name" gorm:"column:name"`
}
