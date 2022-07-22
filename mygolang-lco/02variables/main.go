package main

import "fmt"

const LoginToken string = "iandasdmp" // Public constant

// jwtToken := 3000000
var jwtToken = 3000000

func main() {

	// String

	var username string = "esteban"
	fmt.Println(username)
	fmt.Printf("variables is of type: %T \n", username)

	// Bool //
	// var isLoggedIn bool = "esteban" // true or false
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variables is of type: %T \n", isLoggedIn)

	// Integer //
	var smallVal int = 256 // for most use cases
	// var smallVal uint8 = 256 // for specific use cases
	fmt.Println(smallVal)
	fmt.Printf("variables is of type: %T \n", smallVal)

	// Floating //
	var smallFloat float32 = 255.123125412515212512
	// var smallFloat float64 = 255.123125412515212512
	fmt.Println(smallFloat)
	fmt.Printf("variables is of type: %T \n", smallFloat)

	// Default Value and Aliases //
	var anotherVariable int
	anotherVariable = 3
	fmt.Println(anotherVariable)
	fmt.Printf("variables is of type: %T \n", anotherVariable)

	// Implicit type //
	var website = "google.cl"
	fmt.Println(website)

	// No "var" style -- warlus operation

	bigNumber := 30000
	fmt.Println(bigNumber)

	// Use public staff //
	fmt.Println(LoginToken)
	fmt.Printf("variables is of type: %T \n", LoginToken)

}

// func hello(){
// 	hello := "Hello"
// }
