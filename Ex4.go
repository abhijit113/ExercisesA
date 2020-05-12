package main

import "fmt"

type sen int
var x3 sen
var y int

func main()  {
	fmt.Println(x3)
	fmt.Printf("%T\n",x3)
	x3=42
	fmt.Println(x3)
	y=int(x3) //this is called conversion not casting
	fmt.Println(y)
}