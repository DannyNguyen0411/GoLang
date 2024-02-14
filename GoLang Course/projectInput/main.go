package main

import(
	"fmt"
)

func main(){
	var input string

	fmt.Println("Type some text:")
	fmt.Scanln(&input) 

	fmt.Printf("You entered: %v", input)
}