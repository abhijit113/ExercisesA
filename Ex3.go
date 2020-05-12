package main

import "fmt"

var x2 int
var y2 string
var z2 bool

func main()  {
	x2=42
	y2="James Bond"
	z2=true
	s:=fmt.Sprintf("%v\t %v\t %v\t",x2,y2,z2)
	fmt.Println(s)
}
