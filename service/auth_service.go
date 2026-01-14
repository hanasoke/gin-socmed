package service

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "password not match"}
	}

}
