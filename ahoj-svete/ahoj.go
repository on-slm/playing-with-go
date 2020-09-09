package main

import (
	"fmt"
	"math"
	"math/rand"
)

func multi(x float64, y float64) float64 {
	return x * y
}

func expon(x, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	fmt.Println("Ahoj svete!")

	random := rand.Intn(50)
	fmt.Printf("Cislo %d zije!\n", random)

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(20))

	pi := math.Pi
	fmt.Printf("Delka kruhu o r = 1 je: %g\n", pi)

	xNr := 83275.547952374
	yNr := 547402.503
	fmt.Printf("%g krat %g je %f\n", xNr, yNr, multi(xNr, yNr))

	var x float64 = 2
	var y float64 = 8
	fmt.Printf("%g^%g=%g\n", x, y, expon(x, y))
}
