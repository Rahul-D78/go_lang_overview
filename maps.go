package main

import "fmt"

func main() {

	Stud := make(map[string] int)
	Stud ["array"] = 42
	fmt.Println(Stud["array"])
	fmt.Println(len(Stud))

	
}