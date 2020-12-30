package main

import "fmt"

func main() {

	// var student[10] int;

    // for i := 0; i < 10; i++ {
	// 	student[i] = i+1
	// 	fmt.Println(student[i])
	// }

	student := [5]int{0, 1, 2, 3, 4}

	for i, value := range student {
		fmt.Println(value, i)
	}

}