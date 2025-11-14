package repository

import (
	"rental-book/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRpository(db *gorm.DB) *UserRepository{
	return  &UserRepository{db}
}

func (r *UserRepository) Create(user *entity.User)error{
	return  r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error){
	var user entity.User
	err:=r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}