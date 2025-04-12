package main

import "fmt"

func addNumbers(a int, b int){
	fmt.Println(a + b)
}

func main(){ 
	a := 5
	b := 10
	addNumbers(a, b)
}