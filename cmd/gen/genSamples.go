package main

import (
	"dl-base/pkg/html"
	"dl-base/pkg/sample"
)

func main() {
	html.Render(sample.New(500))
}
