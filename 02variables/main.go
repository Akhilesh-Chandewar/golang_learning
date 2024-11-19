package main

import (
	"fmt"
)

// jwtToken := 3.14

var jwtToken float32 = 3.14

const LoginToken string = "hcgvhggjzxvjh"

func main() {
	var username string = "Akhilesh Chandewar"
	fmt.Println(username)
	fmt.Printf("Variable of type : %T \n" ,  username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of type : %T \n" ,  isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type : %T \n" ,  smallVal)

	var smallFLoat float32 = 255.477
	fmt.Println(smallFLoat)
	fmt.Printf("Variable of type : %T \n" ,  smallFLoat)

	var anotherIntVariable int
	fmt.Println(anotherIntVariable)
	fmt.Printf("Variable of type : %T \n" ,  anotherIntVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable of type : %T \n" ,  anotherStringVariable)

	var website = "https://www.google.com"
	fmt.Println(website)

	website = "https://www.facebook.com"
	fmt.Println(website)

	numberOfPi := 3.14
	fmt.Println(numberOfPi)
	fmt.Printf("Variable of type : %T \n" ,  numberOfPi)

	fmt.Println(jwtToken)
	fmt.Printf("Variable of type : %T \n" ,  jwtToken)

	fmt.Println(LoginToken)
	fmt.Printf("Variable of type : %T \n" ,  LoginToken)
}