package main

import "fmt"

func main1() {
	age := 32
	fmt.Println("Age:", age)         // Print the value of age
	getAdultYears(&age)              // Call the function to get the number of adult years
	fmt.Println("Adult years:", age) // Print the number of adult years
}

func getAdultYears(age *int) {
	*age = *age - 18
}
