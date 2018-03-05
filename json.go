package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	First string `json:"First"`
	Last string	`json:"Last"`
	Age int	`json:"Age"`
}

func main() {
	p1 := Person{"Mr.","Robot",25}
	p2 := Person{"Dr.","X",27}

	people := []Person{p1,p2}
	fmt.Println(people)

	js,err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(js))

	var jsonpeople []Person
	errr := json.Unmarshal(js,&jsonpeople)
	if errr != nil {
		fmt.Println(err)
	}

	for i,v := range jsonpeople{
		fmt.Println("\nPerson",i)
		fmt.Println(v.First,v.Last, v.Age, v)
	}
	

}
