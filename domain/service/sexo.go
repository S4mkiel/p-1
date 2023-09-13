package service

import (
	"github.com/S4mkiel/p-1/domain/entity"
	"github.com/S4mkiel/p-1/domain/repository"
)

type SexoService struct {
	SexoRepository repository.SexoRepository
}


func NewSexoService(sexoRepository repository.SexoRepository) *SexoService {
	return &SexoService{
		SexoRepository: sexoRepository,
	}
}

func (s *SexoService) Create(sx *entity.Sexo) (*entity.Sexo, error) {
	return s.SexoRepository.Create(sx)
}

func (s *SexoService) Update(sx *entity.Sexo) (*entity.Sexo, error) {
	return s.SexoRepository.Update(sx)
}

func (s *SexoService) Get() (*[]entity.Sexo, error) {
	return s.SexoRepository.Get()
}

func (s *SexoService) GetByID(id uint) (*entity.Sexo, error) {
	return s.SexoRepository.GetByID(id)
}