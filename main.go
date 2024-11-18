package main

import (
	"fmt"
	"learning/myutils"
	"learning/variables"
)

func main() {
	fmt.Println("Hello World")
	Demo()
	myutils.PrintMessage("Hello from my utils")
	variables.PrintVariables()
}

func Demo() {
	fmt.Println("Hello from demo")
}