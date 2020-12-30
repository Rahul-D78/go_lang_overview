package main

import "fmt"

func main() {

  x := 10
  
  changeVal(&x)
  fmt.Println(x)
}

func changeVal(x *int)  {
	*x = 7
}