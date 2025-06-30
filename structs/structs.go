package main

import (
	"fmt"

	"example.com/structs/user" // Adjust the import path according to your project structure
)

func main() {
	//  Assligning values to the struct fields directly
	var appUser1 = user.User{
		FirstName: getUserData("Please enter your first name: "),
		LastName:  getUserData("Please enter your last name: "),
		BirthDate: getUserData("Please enter your birth date (YYYY-MM-DD): "),
	}
	fmt.Println(appUser1)
	// Assigning values in order of declaration
	var appUser, err = user.NewUser()
	if err != nil {
		panic(err)
	}

	// Empty struct
	// var emptyUser = user{}
	// fmt.Println("Empty User:", emptyUser)
	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

	var appAdmin, adminErr = user.NewAdmin()
	if adminErr != nil {
		panic(adminErr)
	}
	appAdmin.OutputUserData()

}

func getUserData(prompt string) string {
	var input string
	fmt.Println(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return input
}
