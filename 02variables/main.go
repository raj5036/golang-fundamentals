package main

import "fmt"

func main() {
	var username string = "EngineerDecodes"
	fmt.Println(username)
	fmt.Printf("Type of username is %T & username is %s\n ", username, username)

	var isLoggedIn bool = false
	fmt.Printf("Type of isLoggedIn is %T\n", isLoggedIn)

	var integer int = 2024
	fmt.Printf("Type of integer is %T\n", integer)

	var floatingPointNumber float64 = 2024.234
	fmt.Printf("Type of floatingPointNumber is %T\n", floatingPointNumber)

	const LoginToken string = "XYZ"
	fmt.Println(LoginToken)

	var literal string = `
		something
		is here
		in this string.
	`
	fmt.Println("literal", literal)

	// Declare variables with unknown types
	unknown := "Raj"
	fmt.Println(unknown)
}
