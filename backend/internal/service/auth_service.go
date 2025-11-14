package service

import (
	"errors"
	"rental-book/internal/entity"
	"rental-book/internal/repository"
	"rental-book/pkg/utils"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *AuthService{
	return &AuthService{r}
}

func (s *AuthService) Register(user *entity.User) error{
	exists, _ := s.repo.FindByEmail(user.Email)
	if exists.ID != 0{
		return errors.New("email already registered")
	}

	hash, _ := utils.HashPassword(user.Password)
	user.Password= hash

	return s.repo.Create(user)
}

func (s *AuthService) Login(email, password string)(*entity.User, error){
	user, err:= s.repo.FindByEmail(email)
	if err!=nil{
		return nil, errors.New("Invalid email or password")
	}
	if !utils.CheckPassword(user.Password, password){
		return nil, errors.New("invalid email or passsword")
	}

	return user, nil
}