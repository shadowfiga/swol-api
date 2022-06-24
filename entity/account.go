package entity

type Account struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
	Verified bool   `json:"verified"`
}
