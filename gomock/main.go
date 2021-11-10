package gomock

func main() {

}

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	Create(u *User) (*User, error)
}

type userRepo struct{}

func (r *userRepo) Create(u *User) (*User, error) {
	panic("implement me")
}
