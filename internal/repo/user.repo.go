package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{};
}

// user repo ur
func (ur *UserRepo) GetInfouser() string {
	return "DucThang"
}