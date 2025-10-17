package dbzo

import (
	"errors"
)

type User struct {
	Name        string
	Permissions int
}

func GetUser(name string) (User, error) {
	if name == "admin" {
		user := User{Name: "admin", Permissions: 1}
		return user, nil
	}
	return User{}, errors.New("bad user")
}
