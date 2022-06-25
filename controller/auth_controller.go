package controller

import (
	"encoding/json"
	"net/http"
	"swol-api/entity"
	service "swol-api/service"
)

type AuthController interface {
	Register()
	Login()
	Logout()
	Refresh()
	Verify()
	All(w http.ResponseWriter, r *http.Request)
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

func (*authController) All(w http.ResponseWriter, r *http.Request) {
	accounts := authService.All()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Accounts []entity.Account `json:"accounts"`
	}{
		Accounts: accounts,
	})
}
