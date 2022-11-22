package model

type User struct {
	Id          int
	CompanyName string
	Email       string
	Password    string
	Active      bool
}

func (u *User) EmailIsEmpty() bool {
	return u.Email == ""
}
