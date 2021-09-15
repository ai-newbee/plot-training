package main

import (
	"plot-training/pkg/config"
	"plot-training/pkg/html"
	"plot-training/pkg/sample"
	"plot-training/pkg/vanilla"
)

func main() {
	xyNoBase := sample.New(500, 0.6, 0.0)
	html.Render(xyNoBase, vanilla.Train2LearnScalar(xyNoBase), "plot-learn-scalar.html")

	xyWithBase := sample.New(500, 0.6, 0.2)
	html.Render(xyWithBase, vanilla.Train2learnVector(xyWithBase), "plot-learn-vector.html")

	csvFileName := "3d-scatter-gen.csv"
	xyz := sample.New3DSample(200, csvFileName)

	vanilla.TrainDeepNetwork(xyz)

	html.Render3dClass("3d-class.html", config.DatasetDirName+"/"+csvFileName)
}
