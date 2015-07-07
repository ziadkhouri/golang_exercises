package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
//NOTE: doesn't print first two values in the sequence
func fibonacci() func() int {
	p0, p1 := 0, 1
	return func () int {
		pn := p0 + p1
		p0 = p1
		p1 = pn
		return pn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
