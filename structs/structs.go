package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	// appUser = user{} // null value instance
	appUser := user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	} // Keep a note if you applied the values in correct order

	// ... do something awesome with gathered data!

	outputUserDetails(&appUser)

}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
