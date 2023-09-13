package db

import (
	"github.com/S4mkiel/p-1/domain/entity"
	"gorm.io/gorm"
)

type SexoRepository struct {
	orm *gorm.DB
}

func NewSexoRepository(db *gorm.DB) *SexoRepository {
	return &SexoRepository{orm: db}
}

func (db *SexoRepository) Create(sx *entity.Sexo) (*entity.Sexo, error) {
	result := db.orm.Create(sx)
	if result.RowsAffected == 1 {
		return sx, nil
	} else {
		return nil, result.Error
	}
}

func (db *SexoRepository) Update(sx *entity.Sexo) (*entity.Sexo, error) {
	result := db.orm.Model(&sx).Where("id = ?", sx.ID).Update("sexo", sx.Sexo)
	if result.Error != nil {
		return nil, result.Error
	}
	return sx, nil
}

func (db *SexoRepository) Get() (*[]entity.Sexo, error) {
	var sx []entity.Sexo
	result := db.orm.Find(&sx)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sx, nil
}

func (db *SexoRepository) GetByID(id uint) (*entity.Sexo, error) {
	var sx entity.Sexo

	result := db.orm.Model(&sx).Where("id = ?", id).First(&sx)

	if result.Error != nil {
		return nil, result.Error
	}

	return &sx, nil
}
