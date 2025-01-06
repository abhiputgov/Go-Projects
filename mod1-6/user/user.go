package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("dont have empty strings here")
	}
	return &User{firstName, lastName, birthDate, time.Now()}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password is missing")
	}
	return &Admin{Email: email, Password: password, User: User{FirstName: "ADMIN", LastName: "ADMIN", BirthDate: "---", CreatedAt: time.Now()}}, nil

}

func OutputUserDetails(u *User) {
	fmt.Println((*u).FirstName, (*u).LastName, (*u).BirthDate, (*u).CreatedAt)
}

func (u User) OutputUserDetails2() {
	fmt.Println((u).FirstName, (u).LastName, (u).BirthDate, (u).CreatedAt)
}

func (u *User) ClearUserName() { // this should edit the actual value in the user struct and not the copy of the struct
	(*u).FirstName = ""
	(*u).LastName = ""
}
