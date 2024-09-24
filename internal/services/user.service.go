package services

import "github.com/ducthang001/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service us
func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfouser()
}