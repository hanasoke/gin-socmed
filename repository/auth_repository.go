package repository

import (
	"gin-socmed/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) EmailExist(email string) bool {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}
