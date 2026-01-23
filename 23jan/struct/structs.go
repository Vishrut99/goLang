package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthdate)

	admin:=user.NewAdmin("test@example.com","securepassword")


	admin.User.Output()
	admin.User.ClearUser()
	admin.User.Output()
	
	if err != nil {
		return
	}

	appUser.Output()
	appUser.ClearUser()
	appUser.Output()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
