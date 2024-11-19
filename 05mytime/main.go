package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my time package")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(createdDate)
}
