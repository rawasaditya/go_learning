package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	outPutText("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)
	outPutText("Enter years: ")
	fmt.Scan(&years)
	outPutText("Enter Expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)

	fmt.Printf(`Future value: %.1f 
Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)

}

func outPutText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, years, expectedReturnRate float64) (fv float64, rFv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rFv = fv / math.Pow(1+inflationRate/100, years)
	return
}
