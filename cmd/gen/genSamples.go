package main

import (
	"plot-training/pkg/html"
	"plot-training/pkg/sample"
	"plot-training/pkg/vanilla"
)

func main() {
	html.Render(sample.New(500), vanllia.Train())
}
