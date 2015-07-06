package main

import (
	"fmt"
	"math"
)

const treshhold = 1e-10

func Sqrt(x float64) float64 {
	var z, p float64 = 1.0, 1e10
	for {
		z = z - ((z * z - x) / (z + z))
		if p - z < treshhold {
			break
		}
		p = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
