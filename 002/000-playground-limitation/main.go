package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	// will always return a fixed value of "2009-11-10 23:00:00 UTC"
	fmt.Println("The time is", time.Now())
}
