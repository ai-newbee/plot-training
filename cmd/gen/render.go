package main

import (
	"plot-training/pkg/html"
	"plot-training/pkg/sample"
	vanllia "plot-training/pkg/vanilla"
)

func main() {
	samples := sample.New(500)
	html.Render(samples, vanllia.Train1W(), "plot-learn-scalar.html")
	html.Render(samples, vanllia.Train2W(), "plot-learn-vector.html")
}
