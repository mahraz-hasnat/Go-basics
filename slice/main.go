package main

import "fmt"

func displaySliceInfo(s []int, sliceFrom string){
	fmt.Println("\n\n\n*********** ", sliceFrom ," ***********")
	fmt.Println("slice     : ", s)
	fmt.Println("len slice :  ", len(s))
	fmt.Println("cap slice : ", cap(s))
	fmt.Println("-----------------------------------------------")
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5 ,6}
	slice := arr[1:4]; // 2, 3, 4   ptr : &2 ; len = 3 ; cap : 5
	slice2:= slice[1:2] //3 ptr : &3 ; len = 1 ; cap : 4
	literalSlice := []int{11, 22, 33, 44}
	sliceUsingMake := make([]int, 3, 5)
	
	displaySliceInfo(slice, "Slice From Array")
	displaySliceInfo(slice2, "Slice from Slice")
	displaySliceInfo(literalSlice, "A literal Slice")
	displaySliceInfo(sliceUsingMake, "Slice using make()")
}



/*

Slice can be created in multiple ways

1. From an array
2. from a slice
3. literal slice
4. using default func make()

We can user a slice just like an array..


*/