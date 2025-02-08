package user

import (
	"time"
	"fmt"
	"errors"
)

type User struct {
	FirstName string
	LastName string
	BirthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

// attach struct to function, function turned into method of user struct (receiver)
func (appUser User) OutputUserData(){
	// no need to dereference struct
	// (*appUser).firstName
	fmt.Println(appUser.FirstName, appUser.LastName, appUser.BirthDate)
}

// mutation method of struct, needs pointer to edit that specific version (not edit a copy)
func (appUser *User) ClearUserName(){
	appUser.FirstName = ""
	appUser.LastName = ""
}

// creation/constructor function
func NewUser(firstName string, lastName string, birthDate string) (*User, error){
	if (firstName == "" || lastName == "" || birthDate == ""){
		return nil, errors.New("Must supply params")
	}

	return &User{
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (Admin, error) {
	if email == "" || password == "" {
		return Admin{}, errors.New("Email and password required")
	}

	return Admin{
		email: email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			LastName: "ADMIN",
			BirthDate: "--",
			createdAt: time.Now(),
		},
	}, nil
}