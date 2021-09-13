package main

import (
	"plot-training/pkg/html"
	"plot-training/pkg/sample"
	vanllia "plot-training/pkg/vanilla"
)

func main() {
	xyNoBase := sample.New(500, 0.6, 0.0)
	html.Render(xyNoBase, vanllia.Train2LearnScalar(xyNoBase), "plot-learn-scalar.html")

	xyWithBase := sample.New(500, 0.6, 0.2)
	html.Render(xyWithBase, vanllia.Train2learnVector(xyWithBase), "plot-learn-vector.html")

	html.Render3dClass("3d-class.html")
}
