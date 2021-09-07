package main

import (
	"dl-base/pkg/sample"
	"fmt"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"log"
	"math"
)

const iter = 1000

// https://pkg.go.dev/gorgonia.org/tensor
func prepare(xy sample.XY) (x, y tensor.Tensor) {
	rows := len(xy.X)
	cols := 2
	b := make([]float32, 0, rows*cols)

	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if i == 0 {
				b = append(b, xy.X[j])
			} else {
				b = append(b, 0)
			}
		}
	}
	backX := b
	backY := xy.Y
	// https://pkg.go.dev/gorgonia.org/tensor
	x = tensor.New(tensor.WithShape(len(xy.X), cols), tensor.WithBacking(backX))
	y = tensor.New(tensor.WithShape(len(xy.Y), 1), tensor.WithBacking(backY))
	return
}

func main() {
	xT, yT := prepare(sample.New(500))
	log.Printf("xT :\n%v \n", xT)
	log.Printf("yT :\n%v \n", yT)
	//log.Printf("xT shape:%v \n yT shape:%v \n",xT.Shape(),yT.Shape())

	s := yT.Shape()
	yT.Reshape(s[0])

	//log.Printf("reshape yT shape:%v \n", yT.Shape())
	g := gorgonia.NewGraph()

	X := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))
	Y := gorgonia.NodeFromAny(g, yT, gorgonia.WithName("y"))

	log.Printf("X.Shape:%v \n Y.Shape():%v \n", X.Shape(), Y.Shape())
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
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Printf("Error during iteration: %v: %v\n", i, err)
			log.Fatalln(err)
			break
		}

		if err = solver.Step(model); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("theta: %2.2f  Iter: %v Cost: %2.3f Accuracy: %2.2f \r",
			theta.Value(),
			i+1,
			cost.Value(),
			accuracy(predicted.Data().([]float32), Y.Value().Data().([]float32)))

		machine.Reset() // Reset is necessary in a loop like this
	}
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
