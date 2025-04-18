package main

import "fmt"

func main() {
	fmt.Println("This is main world")
	name := getUserName()
	fmt.Printf("Hello, %s! Welcome to the world of Go.\n", name)
	// add(5, 10)	

	func(a, b int) { // anonymous function
		fmt.Println(a+b)
	}(5,7) // imidiately invoked function expression (IIFE)

	multiply := func(a, b int) int { // anonymous function
		return a * b
	}

	fmt.Println(multiply(5, 10)) // invoking the anonymous function multiply
}

// The init function is a special function in Go that is executed before the main function.
// It is used to initialize variables, set up resources, or perform any setup tasks that need to be done before the program starts running.
//* The init function is executed before the main function in Go.

func init() {
	fmt.Println("This is init world, and it is executed before main")
}

// standard or named function
// The getUserName function prompts the user for their name and returns it as a string.
func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")	
	fmt.Scanln(&name)
	return name
}

func add(a, b int){ // another standard function
	fmt.Println(a + b)
}