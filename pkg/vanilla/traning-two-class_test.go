package vanilla

import (
	"plot-training/pkg/sample"
	"testing"
)

func init() {

}
func TestTrainDeepNetwork(t *testing.T) {
	csvFileName := "3d-scatter-gen.csv"
	xyz := sample.New3DSample(200, csvFileName)

	ret := TrainDeepNetwork(xyz)
	if "ok" != ret {
		t.Fatal("err")
	}
}
