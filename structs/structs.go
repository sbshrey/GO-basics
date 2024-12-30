package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser *user.User
	// appUser = user{} // null value instance

	admin, err := user.NewAdmin("test@example.com", "test123")

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser, err := user.New(userFirstName, userLastName, userBirthdate) // Keep a note if you applied the values in correct order

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
