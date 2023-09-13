package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username,omitempty" gorm:"type:varchar"`
	Password   string `json:"password,omitempty" gorm:"type:varchar"`
	Name       string `json:"name,omitempty" gorm:"type:varchar"`
	Telefone   string `json:"telefone,omitempty" gorm:"type:varchar"`
	Email      string `json:"email,omitempty" gorm:"type:varchar"`
	SexoId     uint   `json:"sexo_id,omitempty" gorm:"type:uint"` // 1 - Masculino, 2 - Feminino, 3 - Outros
	Sexo       Sexo   `json:"-" gorm:"foreignKey:SexoId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Nascimento string `json:"nascimento,omitempty" gorm:"type:varchar"`
}
