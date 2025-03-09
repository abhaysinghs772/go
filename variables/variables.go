package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println("a = ", a)

	var b int = 1
	var c int = 2
	// var b, c int = 1, 2 // this is also valid
	fmt.Println("b + c = ", b+c)

	var d = true
	fmt.Println("d = ", d)

	var e int
	fmt.Println(e) // print 0

	f := "apple" // this is also valid
	fmt.Println(f)
}
