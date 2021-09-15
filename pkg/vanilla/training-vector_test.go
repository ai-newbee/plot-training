package vanilla

import (
	"plot-training/pkg/sample"
	"testing"
)

func TestTrain2learnVector(t *testing.T) {
	xyWithBase := sample.New(5000, 0.6, 0.2)
	Train2learnVector(xyWithBase)
}
