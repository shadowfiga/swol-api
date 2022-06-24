package controller

import service "swol-api/service"

type AuthController interface {
	Register()
	Login()
	Logout()
	Refresh()
	Verify()
}

var (
	authService service.AuthService = service.NewAuthService()
)

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (*authController) Register() {

}

func (*authController) Login() {

}

func (*authController) Logout() {

}

func (*authController) Refresh() {

}

func (*authController) Verify() {

}
