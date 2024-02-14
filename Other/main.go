package main

import "fmt"

func main(){
	var input string

	fmt.Println("Mein Kampf")
	fmt.Scanln(&input)

	fmt.Println("You entered: ", input)
}