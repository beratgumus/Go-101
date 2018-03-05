package main

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9}
	s := sum(ii...)
	println(s)

	s2 := even(sum,ii...)
	println(s2)

}

func sum(slice ...int) int {
	total :=0
	for _,v := range slice{
		total += v
	}
	return total
}

func even(f func(newSlice ...int) int , slice ...int) int {  //func as a parameter
	var evenslice []int
	for _,v := range slice {
		if v%2 == 0{
			evenslice = append(evenslice,v)
		}
	}
	return f(evenslice...)
}