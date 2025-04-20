package main

import "fmt"

func print(nums *[3]int){ // nums is a pointer to an array of 3 integers
	fmt.Println(nums) // 0xc00000c0a0
}

func main() {
	arr := [3]int{1,2,3}
	print(&arr) // pass the address of arr to the function
}


//  a := 1
//  addOfA := &a // &a is the address of a
//  fmt.Println(a)    // 1
//  fmt.Println("Address of a = ", addOfA) // 1

//  *addOfA = 2 // dereference the pointer to change the value of a
//  fmt.Println("After updating a through pointer", a)    // 2