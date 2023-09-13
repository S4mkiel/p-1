package repository

import "github.com/S4mkiel/p-1/domain/entity"

type UserRepository interface {
	Create(*entity.User) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
	Get() (*[]entity.User, error)
	GetByID(id uint) (*entity.User, error)
}
