package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	FindById(id int) (*User, error)
	Save(user *User) error
}
