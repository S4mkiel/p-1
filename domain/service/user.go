package service

import (
	"github.com/S4mkiel/p-1/domain/entity"
	"github.com/S4mkiel/p-1/domain/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}


func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}


func (s *UserService) Create(u *entity.User) (*entity.User, error) {
	return s.UserRepository.Create(u)
}

func (s *UserService) Update(u *entity.User) (*entity.User, error) {
	return s.UserRepository.Update(u)
}

func (s *UserService) Get() (*[]entity.User, error) {
	return s.UserRepository.Get()
}

func (s *UserService) GetByID(id uint) (*entity.User, error) {
	return s.UserRepository.GetByID(id)
}