package main

import "fmt"

func main() {
	//  Add comments to explain the code
	// Declare an integer variable a and assign it a value of 42
	var a int = 42 // This is an integer variable
	// Declare a pointer variable p that points to the address of a
	var p *int = &a                                                     // This is a pointer variable that holds the address of a
	fmt.Println("Value of a:", a)                                       // Print the value of a
	fmt.Println("Address of a:", &a)                                    // Print the address of a
	fmt.Println("Value of p:", p)                                       // Print the value of p, which is the address of a
	fmt.Println("Value pointed to by p:", *p)                           // Print the value pointed to by p, which is the value of a
	*p = 10                                                             // Modify the value of a through the pointer p
	fmt.Println("New value of a after modifying through pointer p:", a) // Print the new value of a after modification
	fmt.Println("Value pointed to by p after modification:", *p)        // Print the value pointed

}
