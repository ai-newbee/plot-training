package main

import (
	"fmt"
	. "gorgonia.org/gorgonia"
	vanllia "plot-training/pkg/vanilla"
)

// Linear Regression Example
//
// The formula for a straight line is
//		y = mx + c
// We want to find an `m` and a `c` that fits the equation well. We'll do it in both float32 and float64 to showcase the extensibility of Gorgonia
func main() {
	var m Value
	// Float32
	m = vanllia.Train2LearnScalar(Float32, 1000)
	fmt.Printf("float32: y = %3.3fx \n", m)
}
