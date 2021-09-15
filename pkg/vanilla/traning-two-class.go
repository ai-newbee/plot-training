package vanilla

import (
	"fmt"
	g "gorgonia.org/gorgonia"
	t "gorgonia.org/tensor"
	"plot-training/pkg/sample"
)

func TrainDeepNetwork(xyz sample.XYZ) string {
	graph := g.NewGraph()
	xy1, z := prepareSample(graph, xyz)
	pre, cost, mat, vec, machine := setupGraph(graph, xy1, z)
	fmt.Println(pre)
	defer machine.Close()

	iter := 1000
	var err error

	models := []g.ValueGrad{mat, vec}

	fmt.Println(mat)

	solver := g.NewVanillaSolver(g.WithLearnRate(0.001))
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Printf("Error during iteration: %v: %v\n", i, err)
			panic(err)
		}

		if err = solver.Step(models); err != nil {
			panic(err)
		}

		//if (i+1)%recordeStripe == 0 {
		//	records = append(records, LostAndW{cost.Value().Data().(float32), m.Value().Data().(float32), 0})
		//}
		//if (i + 1) == iter {
		//	fmt.Printf("m: %v  iter: %v Cost: %2.3f  \n",
		//		m.Value(),
		//		i+1,
		//		cost.Value())
		//
		//	fmt.Println(records)
		//}
		fmt.Printf("cost: %f \n", cost.Value().Data().(float32))

		machine.Reset() // Reset is necessary in a loop like this
	}
	return "ok"
}

//extra one column for learning b (basis)
func prepareSample(graph *g.ExprGraph, xyz sample.XYZ) (xy1, z *g.Node) { //z is ground-truth
	sampleCount := len(xyz.X)
	back := make([]float32, 0, sampleCount*3)
	for i := 0; i < sampleCount; i++ {
		back = append(back, xyz.X[i], xyz.Y[i], 1)
	}
	xy1Tensor := t.New(t.WithShape(sampleCount, 3), t.WithBacking(back))
	zTensor := t.New(t.WithShape(sampleCount), t.WithBacking(xyz.Z))

	s := zTensor.Shape()
	zTensor.Reshape(s[0])

	xy1 = g.NodeFromAny(graph, xy1Tensor, g.WithName("xy1"))

	z = g.NodeFromAny(graph, zTensor, g.WithName("z"))
	return
}

func setupGraph(graph *g.ExprGraph, xy1, z *g.Node) (predict, cost, mat, vec *g.Node, machine g.VM) {
	mat = g.NewMatrix(
		graph,
		g.Float32,
		g.WithName("mat"),
		g.WithShape(3, 3), // x,y,1 dim is 3
		g.WithInit(g.Ones()))

	vec = g.NewVector(
		graph,
		g.Float32,
		g.WithName("vec"),
		g.WithShape(3), // x,y,1 dim is 3
		g.WithInit(g.Ones()))
	hidden := g.Must(g.Sigmoid(g.Must(g.Mul(xy1, mat))))
	predict = g.Must(g.Mul(hidden, vec))

	squaredError := must(g.Square(must(g.Sub(predict, z))))

	cost = must(g.Mean(squaredError))

	//!!! must do it before new a machine
	if _, err := g.Grad(cost, mat, vec); err != nil {
		panic(err)
	}
	machine = g.NewTapeMachine(graph, g.BindDualValues(mat, vec))
	return
}
