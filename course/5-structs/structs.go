package main

import (
	"fmt"
	"structs/user"
)

// type user struct {
// 	firstName string
// 	lastName string
// 	birthDate string
// 	createdAt time.Time
// }

type customString string

func main(){
	userFirstName := getUserInput("Please enter first name: ")
	userLastName := getUserInput("Please enter last name: ")
	userBirthDate := getUserInput("Please enter DOB: ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if (err != nil){
		fmt.Println(err)
		return
	}

	emailInput := getUserInput("Enter admin email: ")
	passwordInput := getUserInput("Password input: ")

	admin, err := user.NewAdmin(emailInput, passwordInput)
	if (err != nil){
		fmt.Println(err)
		return
	}
	admin.User.OutputUserData()

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

	var var1 customString = "Josh"
	var1.log()
}

func (text customString) log(){
	fmt.Println(text)
}

func getUserInput(input string) string{
	fmt.Print(input)
	var value string
	fmt.Scanln(&value)
	return value
}

