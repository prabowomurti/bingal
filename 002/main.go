package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")

	// Go also has a built-in function from the runtime
	print("Hello, World\n")
	println(27, "my age")

	// but it fails to print a more complex datatype
	// println(time.Now())

}
