package db

import (
	"github.com/S4mkiel/p-1/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	orm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{orm: db}
}

func (db *UserRepository) Create(u *entity.User) (*entity.User, error) {
	result := db.orm.Create(u)
	if result.RowsAffected == 1 {
		return u, nil
	} else {
		return nil, result.Error
	}
}

func (db *UserRepository) Update(u *entity.User) (*entity.User, error) {
	result := db.orm.Model(&u).Where("id = ?", u.ID).Updates(entity.User{Email: u.Email, Password: u.Password, SexoId: u.SexoId, Telefone: u.Telefone})
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (db *UserRepository) Get() (*[]entity.User, error) {
	var user []entity.User
	result := db.orm.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *UserRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	
	result := db.orm.Model(&user).Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
