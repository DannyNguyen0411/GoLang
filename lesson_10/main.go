package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	} 

	// ints as key type
	phonebook := map[int]string{
		112: "Police",
		113: "Vietnamese Police",
		119: "Korean Police",
		911: "American Police",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[112])

	phonebook[113] = "Suicide Prevention"
	fmt.Println(phonebook)

	phonebook[112] = "Suicide Terrorist"
	fmt.Println(phonebook)
	

}