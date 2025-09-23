package service

import (
	"errors"
	"umkm-api/internal/user/model"
	"umkm-api/internal/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(name, email, password string) (*model.User, error)
	Login(email, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(name, email, password string) (*model.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{Username: name, Email: email, Password: string(hashed)}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("password salah")
	}
	return user, nil
}
