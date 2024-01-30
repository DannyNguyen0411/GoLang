package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 { 
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 40")
	}

	names := []string{"adolf hitler", "genghis khan", "napoleon", "ash kethcum", "toad"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2{
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at position %v is %v\n", index, value)
	}
}