package repository

type AuthRepository interface {
	Create()
	GetById()
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