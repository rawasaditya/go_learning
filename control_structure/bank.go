package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const acntBalancefile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(acntBalancefile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (acntBalance float64, err error) {
	data, err := os.ReadFile(acntBalancefile)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	balanceText := string(data)
	acntBalance, err = strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return
}

func main() {

	accountBalance, err := getBalanceFromFile()
	if err != nil {
		panic(err)
	}
	for {
		// Display choice
		fmt.Println("\nWelcome to Go Bank!")
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		// Get choice
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposite ?")
			var deposite float64
			fmt.Scan(&deposite)
			if deposite <= 0 {
				fmt.Println("Invalid amount. Must be greater that zero")
				continue
			}
			accountBalance += deposite
			writeBalanceToFile(accountBalance)
			fmt.Println("Your account balance is updated to ", accountBalance)
		case 3:
			fmt.Println("How much do you want to withDraw ?")
			var withDraw float64
			fmt.Scan(&withDraw)
			if withDraw <= 0 {
				fmt.Println("Invalid amount. Must be greater that zero")
				continue
			}

			if withDraw > accountBalance {
				fmt.Println("Invalid amount. Must be within", accountBalance)
				continue
			}
			accountBalance -= withDraw
			writeBalanceToFile(accountBalance)
			fmt.Println("Your account balance is updated to ", accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for chosing our bank")
			return
		}
	}
}
