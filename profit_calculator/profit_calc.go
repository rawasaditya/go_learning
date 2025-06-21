package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	getUserValue("Enter your revenue: ", &revenue)
	getUserValue("Enter your expenses: ", &expenses)
	getUserValue("Enter your tax rate: ", &taxRate)
	ebt, profit, ration := calculateEbtProfitRation(revenue, expenses, taxRate)
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ration)

}

func getUserValue(text string, toExcept *float64) {
	fmt.Print(text)
	fmt.Scan(toExcept)
}

func calculateEbtProfitRation(revenue, expenses, taxRate float64) (ebt float64, profit float64, ration float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ration = ebt / profit
	return
}
