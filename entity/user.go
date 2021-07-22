package entity

type User struct {
	ID        ID
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func NewUser(firstName, lastName, email, phone string) (*User, error) {
	return &User{
		ID:        NewID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}, nil
}

func (u *User) Validate() bool {
	return u.FirstName == "" || u.Email == ""
}
