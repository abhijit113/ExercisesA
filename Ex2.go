package main

import "fmt"

var x1 int // x should not be declared in the same package
var y1 string
var z1 bool

func main()  {
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)

	fmt.Println("These default values are called zero values")
}
