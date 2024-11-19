package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("Value of pointer is: ", &ptr)

	myNumber := 23
	var myNumberPtr *int = &myNumber
	fmt.Println("Value of myNumberPtr is: ", myNumberPtr)
	fmt.Println("Value of myNumberPtr is: ", *myNumberPtr)
}
