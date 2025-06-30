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
	createAt  time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputUserData() {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", u.FirstName, u.LastName, u.BirthDate)
}

func (a Admin) OutputUserData() {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\nEmail: %s\nPassword: %s\nCreated At: %s", a.FirstName, a.LastName, a.BirthDate, a.email, a.password, a.createAt)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
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
func NewUser() (*User, error) {
	var firstName = getUserData("Please enter your first name: ")
	var lastName = getUserData("Please enter your last name: ")
	var birthDate = getUserData("Please enter your birth date (YYYY-MM-DD): ")
	var createAt = time.Now()
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		createAt:  createAt,
	}, nil
}

func NewAdmin() (*Admin, error) {
	user, err := NewUser()
	if err != nil {
		return nil, err
	}
	var email = getUserData("Please enter your email: ")
	var password = getUserData("Please enter your password: ")
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil

}
