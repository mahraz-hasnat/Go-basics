package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	user1 := User{ID: 1, Name: "Alice", Age: 30}
	user2 := User{ID: 2, Name: "Bob", Age: 25}

	fmt.Println("ID of user1:", user1.ID)
	fmt.Println("Name of user1:", user1.Name)
	fmt.Println("Age of user1:", user1.Age)
	fmt.Println("ID of user2:", user2.ID)
	fmt.Println("Name of user2:", user2.Name)
	fmt.Println("Age of user2:", user2.Age)
}
