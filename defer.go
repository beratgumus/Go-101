package main

import "fmt"

func main() {
	// Defer
	fmt.Println("\nDefer:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		// Deferred things are added to a stack and executed after func return
		defer fmt.Printf("%d ", i)
	}
	fmt.Println()

	defer first()
	second()
}

func first(){
	fmt.Println("this is func first")
}

func second(){
	fmt.Println("this is func second")
}
