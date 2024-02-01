package main

import (
	"fmt"
	"strings"
)

func genInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names{
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main(){
	fn1, sn1 := genInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := genInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := genInitials("barret")
	fmt.Println(fn3, sn3)
}