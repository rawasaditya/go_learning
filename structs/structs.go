package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

func main() {
	//  Assligning values to the struct fields directly
	// var appUser = user{
	// 	firstName: getUserData("Please enter your first name: "),
	// 	lastName:  getUserData("Please enter your last name: "),
	// 	birthDate: getUserData("Please enter your birth date (YYYY-MM-DD): "),
	// 	createAt:  time.Now(),
	// }

	// Assigning values in order of declaration
	var appUser = user{
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birth date (YYYY-MM-DD): "),
		time.Now(),
	}

	// Empty struct
	// var emptyUser = user{}
	// fmt.Println("Empty User:", emptyUser)

	outputUserData(&appUser)

}

func outputUserData(userValues *user) {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", userValues.firstName, userValues.lastName, userValues.birthDate)
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
