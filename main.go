package main

import "fmt"


func main() {
	fmt.Println("This is main world")
	name := getUserName()
	fmt.Printf("Hello, %s! Welcome to the world of Go.\n", name)
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