package service

import (
	"swol-api/entity"
	"swol-api/repository"
)

type AuthService interface {
	Register()
	Login()
	Logout()
	Refresh()
	Verify()
	All() []entity.Account
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

var (
	authRepository = repository.NewAuthRepository()
)

func (*authService) Register() {

}

func (*authService) Login() {

}

func (*authService) Logout() {

}

func (*authService) Refresh() {

}

func (*authService) Verify() {

}

func (*authService) All() []entity.Account {
	return authRepository.GetAll()
}
