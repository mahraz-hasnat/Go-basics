package main

import "fmt"

func main() {
	arr := getArrayLength()
	getArrayValues(arr)
	displayArray(arr)
}

func getArrayLength() []int {
	var arrLen int
	fmt.Print("Enter the length of the array: ")
	fmt.Scan(&arrLen)
	arr := make([]int, arrLen)
	return arr
}

func getArrayValues(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print("Enter value for element ", i, ": ")
		fmt.Scan(&arr[i])
	}
}

func displayArray(arr []int) {
	fmt.Println("The array is:", arr)
	fmt.Println("The length of the array is:", len(arr))
	fmt.Println("The capacity of the array is:", cap(arr))
	fmt.Println("The address of the array is:", &arr)
	fmt.Println("The address of the first element is:", &arr[0])
	fmt.Println("---------------------------------------------")
}