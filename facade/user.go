package facade

type User struct {
	email string
}

func NewUser(email string) *User {
	return &User{
		email: email,
	}
}