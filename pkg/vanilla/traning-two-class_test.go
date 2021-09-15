package vanilla

import (
	"plot-training/pkg/sample"
	"testing"
)

func init() {

}
func TestTrainDeepNetwork(t *testing.T) {
	csvFileName := "3d-scatter-gen.csv"
	xyz := sample.New3DSample(2000, csvFileName)

	ret := TrainDeepNetwork(xyz)
	if "ok" != ret {
		t.Fatal("err")
	}
}
