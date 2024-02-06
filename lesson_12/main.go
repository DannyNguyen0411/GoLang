package main

import (
	"fmt"
)

// Function that accept pointers
func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name)

	memory_address := &name
	// fmt.Println("memory address: ", memory_address)

	// // Astrics * will point out the value of the memory that is pointed out.
	// fmt.Println("value at memory address:", *memory_address)

	fmt.Println(name)
	updateName(memory_address)
	fmt.Println(name)

	// fmt.Println(name)
}
