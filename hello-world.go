package main

import (
	"fmt"
)

// This type of declaration is not allowed in the package scope
//c:= 3.14

// allowed in the package scope
var d int = 5 
var e = 3.14 
var f int
var a, b, c int = 1, 3, 5

var (
    myNum int = 10
    //newNum := 20
    // var myString string = "Hello, World!"
    myString = "Hello, World!"
    myFloat float64 = 3.14
)

func main() {
    var a string;
    a = "Learning GO is fun!"
    b := 42 // this is allowed in the function scope. also we must put a value to it.   
    i,b := "--",100 // this is allowed in the function scope. also we must put a value to it.

    randomTexts := "This is a random text"

    fmt.Println(randomTexts, " ", myNum, " ", myString, " ", myFloat, " ", a, " ", b, " ", i, " ", c, " ", d, " ", e, " ", f)
    // fmt.Println(a)
    // fmt.Println(b)
    // fmt.Println(i)

}