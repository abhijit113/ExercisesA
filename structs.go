package main

import "fmt"

type example struct {
	flag    bool
	counter int
	pi      float32
}

func main() {
	var e1 example
	fmt.Printf("%+v\n", e1)
	fmt.Printf("%#v\n", e1)

	e2 := example{}
	fmt.Printf("%#v\n", e2)

	e2 = example{
		flag:    true,
		counter: 10,
		pi:      3.467,
	}

	fmt.Printf("%+v\n", e2)

}
