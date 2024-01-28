package main

import (
	"fmt"
	"sort"
)

func main() {
	// greetings := "Greetings, I'm Chara"
	// Strings
	// fmt.Println(strings.Contains(greetings, "Chara"))
	// fmt.Println(strings.ReplaceAll(greetings, "Chara", "an Asian"))
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "a"))
	// fmt.Println(strings.Split(greetings, " "))

	// the original value is unchanged
	// fmt.Println("original string value =", greetings)

	// Sort in the order
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"adolf hitler", "genghis khan", "napoleon", "ash kethcum", "toad"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "toad"))
}
