package vanilla

import (
	"fmt"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"log"
	"math"
	"plot-training/pkg/config"
	"plot-training/pkg/sample"
)

// https://pkg.go.dev/gorgonia.org/tensor
func prepare(xy sample.XY) (x, y tensor.Tensor) {
	rows := len(xy.X)
	cols := 2
	x1x0 := make([]float32, 0, rows*cols)
	for j := 0; j < rows; j++ {
		x1x0 = append(x1x0, xy.X[j], 1)
	}
	backX := x1x0
	backY := xy.Y
	// https://pkg.go.dev/gorgonia.org/tensor
	x = tensor.New(tensor.WithShape(len(xy.X), cols), tensor.WithBacking(backX))
	y = tensor.New(tensor.WithShape(len(xy.Y), 1), tensor.WithBacking(backY))
	return
}

func Train2learnVector(xy sample.XY) []LostAndW {
	xT, yT := prepare(xy)
	log.Printf("xT :\n%v \n", xT)
	log.Printf("yT :\n%v \n", yT)
	log.Printf("xT :%v \n yT :%v \n", xT, yT)

	s := yT.Shape()
	yT.Reshape(s[0])

	//log.Printf("reshape yT shape:%v \n", yT.Shape())
	g := gorgonia.NewGraph()

	X := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))

	Y := gorgonia.NodeFromAny(g, yT, gorgonia.WithName("y"))

	log.Printf("X.Shape:%v \n Z.Shape():%v \n", X.Shape(), Y.Shape())
	//gorgonia.Gaussian(0, 0) means no any random chance , but gorgonia.Zeroes() is better
	theta := gorgonia.NewVector(
		g,
		gorgonia.Float32,
		gorgonia.WithName("theta"),
		gorgonia.WithShape(X.Shape()[1]),
		gorgonia.WithInit(gorgonia.Zeroes()))

	//theta := gorgonia.NewScalar(g, gorgonia.Float32, gorgonia.WithName("theta"))
	//gorgonia.Let(theta, 0.1)
	log.Printf("theta.Shape:%v \n ", theta.Shape())

	pred := must(gorgonia.Mul(X, theta))

	// Gorgonia might delete values from nodes so we are going to save it
	// and print it out later
	var predicted gorgonia.Value
	//.Read is reading node into value
	gorgonia.Read(pred, &predicted)

	squaredError := must(gorgonia.Square(must(gorgonia.Sub(pred, Y))))

	//define cost as mean squared error
	cost := must(gorgonia.Mean(squaredError))

	if _, err := gorgonia.Grad(cost, theta); err != nil {
		//log.Fatalf("Failed to backpropagate: %v", err)
		panic(err)
	}

	machine := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(theta))
	defer machine.Close()

	//define grad descent as optimiser
	model := []gorgonia.ValueGrad{theta}

	//learning rate , solve method, detailed how
	solver := gorgonia.NewVanillaSolver(gorgonia.WithLearnRate(0.01))

	var err error
	const iter = config.IterVector
	const recordeStripe = config.RecordeStripe4Vector
	records := make([]LostAndW, 0, iter/recordeStripe)
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Printf("Error during iteration: %v: %v\n", i, err)
			log.Fatalln(err)
			break
		}

		if err = solver.Step(model); err != nil {
			log.Fatal(err)
		}

		if (i+1)%recordeStripe == 0 {
			data := theta.Value().Data().([]float32)
			records = append(records, LostAndW{cost.Value().Data().(float32), data[0], data[1]})
		}
		if (i + 1) == iter {
			fmt.Printf("theta: %v  iter: %v Cost: %2.3f Accuracy: %2.2f \n",
				theta.Value(),
				i+1,
				cost.Value(),
				accuracy(predicted.Data().([]float32), Y.Value().Data().([]float32)))

			fmt.Println(records)
		}
		machine.Reset() // Reset is necessary in a loop like this
	}
	return records
}

func accuracy(prediction, y []float32) float32 {
	var ok float32
	for i := 0; i < len(prediction); i++ {
		if math.Round(float64(prediction[i]-y[i])) == 0 {
			ok += 1.0
		}
	}
	return ok / float32(len(y))
}

func must(n *gorgonia.Node, err error) *gorgonia.Node {
	if err != nil {
		panic(err)
	}
	return n
}
