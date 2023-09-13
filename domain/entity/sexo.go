package entity

import "gorm.io/gorm"

type Sexo struct {
	gorm.Model
	Sexo string `json:"sexo,omitempty" gorm:"type:varchar"`
}