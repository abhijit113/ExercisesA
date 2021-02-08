package main

import "fmt"

type example struct {
	flag    bool
	counter int
	pi      float32
}

type apple struct {
	flag    bool
	counter int
	pi      float32
}

func main() {

	e1 := example{
		flag:    true,
		counter: 50,
		pi:      3.467,
	}

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.467,
	}

	a2 := apple{
		flag:    true,
		counter: 10,
		pi:      3.467,
	}

	e1 = e2
	//e1,e2,a2 are name type

	literal := struct {
		flag    bool
		counter int
		pi      float32
	}{
		flag:    true,
		counter: 500,
		pi:      3.467,
	}
	a2 = literal

	fmt.Printf("%+v - %+v - %+v\n", e1, e2, a2)

}
