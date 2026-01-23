package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User     //Embedding User struct
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		User: User{ //Initializing embedded User struct
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "01/01/1970",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}

func (u *User) Output() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUser() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}
