package service

type AuthService interface {
	Register()
	Login()
	Logout()
	Refresh()
	Verify()
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

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
