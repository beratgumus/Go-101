package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nSlices:")
	slice := make([]int,0)

	for i:= 0; i<5; i++ {
		slice = append(slice,i)
	}
	fmt.Println(slice)

	slice2 := make([]int,2)
	slice = append(slice,slice2...)
	fmt.Println(slice)

	fmt.Println("\nMaps:")

	strMap := map[string]func(string)string{
		"append1": func(el string) string {
			return el + "1"
		},
		"uppercase": func(el string) string {
			return strings.ToUpper(el)
		},
	}

	fmt.Printf("el1: %s\n", strMap["uppercase"]("wow"))
	fmt.Printf("el2: %s", strMap["append1"]("wow"))
}
