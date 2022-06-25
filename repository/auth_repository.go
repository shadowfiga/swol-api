package repository

import (
	"context"
	"swol-api/db"
	"swol-api/entity"
)

type AuthRepository interface {
	Create()
	GetById()
	GetAll() []entity.Account
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (*authRepository) Create() {

}

func (*authRepository) GetById() {

}

func (*authRepository) GetAll() []entity.Account {
	cursor, _ := db.Db.Collection("accounts").Find(context.TODO(), nil)
	accounts := make([]entity.Account, 0)
	cursor.All(context.TODO(), accounts)
	return accounts
}
