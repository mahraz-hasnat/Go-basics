package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func (user User) displayDetails() { // receiver function
	fmt.Println("User ID:", user.ID)
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Age:", user.Age)
}

func main() {
	user1 := User{ID: 1, Name: "Alice", Age: 30}
	user2 := User{ID: 2, Name: "Bob", Age: 25}

	// printUser(user1)	
	// printUser(user2)

	user1.displayDetails()
	user2.displayDetails()
}
