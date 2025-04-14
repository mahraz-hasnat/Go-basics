package main

import (
	"fmt"
)

func main() {
    /*
    var myString string = "learing about GO!"
    var myInt int = 12

    ////   PrnitF is a function that formats and prints its arguments to the standard output.

    // The %v verb is used to print the value in a default format.
    fmt.Printf("%v\n", myInt)
    fmt.Printf("%v\n", myString)

    // The %#v verb is used to print the value in Go syntax.
    fmt.Printf("%#v\n", myInt)
    fmt.Printf("%#v\n", myString)

    // The %T verb is used to print the type of the value.
    fmt.Printf("%T\n", myInt)
    fmt.Printf("%T\n", myString)
*/

    type myType int
    var myVar myType = 42
    fmt.Printf("%T\n", myVar)
}