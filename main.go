package main

import "fmt"


func main() {
	fmt.Println("This is main world")
}

// The init function is a special function in Go that is executed before the main function.
// It is used to initialize variables, set up resources, or perform any setup tasks that need to be done before the program starts running.
//* The init function is executed before the main function in Go.

func init() {
	fmt.Println("This is init world, and it is executed before main")
}