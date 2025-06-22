package main

import (
	"fmt"

	"example.com/bank/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const acntBalancefile = "balance.txt"

func main() {

	accountBalance, err := fileOps.GetFloatFromFile(acntBalancefile)
	if err != nil {
		panic(err)
	}
	for {
		presentOptions()
		// Get choice
		var choice int
		fmt.Println(randomdata.SillyName())

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
			fileOps.WriteFloatToFile(acntBalancefile, accountBalance)
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
			fileOps.WriteFloatToFile(acntBalancefile, accountBalance)
			fmt.Println("Your account balance is updated to ", accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for chosing our bank")
			return
		}
	}
}
