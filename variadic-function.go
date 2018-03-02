package main

import "fmt"

func main() {

	fmt.Println(sum(1,2,3,4,5))

	arr := []int{1,2,3,4,5}
	fmt.Println(sum(arr...))
	}

func sum(list ...int) (sum int){
	sum = 0
	for _,v := range list{
		sum += v
	}
	return
}