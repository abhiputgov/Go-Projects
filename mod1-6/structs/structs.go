package main

import (
	"fmt"

	"example.com/b/user"
)

func main() {
	firstName := getUserData("FirstName:")
	lastName := getUserData("LastName:")
	birthDate := getUserData("Birthdate in mm/dd/yyyy format:")
	user1, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin1, err := user.NewAdmin("test@example.com", "test123")

	if err != nil {
		fmt.Println(err)
		return
	}

	admin1.OutputUserDetails2()
	admin1.ClearUserName()
	admin1.OutputUserDetails2()

	user.OutputUserDetails(user1)
	user1.OutputUserDetails2()
	user1.ClearUserName()
	user1.OutputUserDetails2()
}

func getUserData(query string) string {
	fmt.Print(query)
	var value string
	fmt.Scanln(&value)
	return value
}
