package repository

import "github.com/S4mkiel/p-1/domain/entity"

type SexoRepository interface {
	Create(*entity.Sexo) (*entity.Sexo, error)
	Update(*entity.Sexo) (*entity.Sexo, error)
	Get() (*[]entity.Sexo, error)
	GetByID(id uint) (*entity.Sexo, error)
}
