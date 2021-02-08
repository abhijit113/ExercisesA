package main

import "fmt"

func main() {
	const (
		A = iota
		B = iota
		C = iota
	)

	fmt.Println(A, B, C)

	const (
		A1 = iota + 1
		B1
		C1
	)

	fmt.Println(A1, B1, C1)

}
