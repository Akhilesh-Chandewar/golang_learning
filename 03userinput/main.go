package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)
	fmt.Printf("Type of rating is %T \n", rating)

	fmt.Println("")

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	fmt.Printf("Type of name is %T \n", name)

	fmt.Println("")

	buf := make([]byte, 10)
	fmt.Println("Enter your name: ")
	fmt.Scan(&buf)
	fmt.Println("Hello", string(buf))
	fmt.Printf("Type of name is %T", buf)
}
