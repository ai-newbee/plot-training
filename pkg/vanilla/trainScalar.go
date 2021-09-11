package vanllia

import (
	"fmt"
	"log"
	"plot-training/pkg/config"
	"plot-training/pkg/sample"
	"runtime"

	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// manually generate a fake dataset which is y=2x+random
func xy(sample sample.XY) (x tensor.Tensor, y tensor.Tensor) {
	const count = 500
	xBack := sample.X
	yBack := sample.Y

	x = tensor.New(tensor.WithBacking(xBack), tensor.WithShape(count))
	y = tensor.New(tensor.WithBacking(yBack), tensor.WithShape(count))
	return
}

func linregSetup(sample sample.XY) (m, cost *Node, machine VM) {
	var xT, yT Value
	xT, yT = xy(sample)

	g := NewGraph()
	x := NewVector(g, Float32, WithShape(xT.Shape()[0]), WithName("x"), WithValue(xT))
	y := NewVector(g, Float32, WithShape(xT.Shape()[0]), WithName("y"), WithValue(yT))
	m = NewScalar(g, Float32, WithName("m"), WithValue(float32(0)))
	//c = NewScalar(g, Float, WithName("c"), WithValue(random(Float)))

	pred := Must(Mul(x, m))
	se := Must(Square(Must(Sub(pred, y))))
	cost = Must(Mean(se))

	if _, err := Grad(cost, m); err != nil {
		log.Fatalf("Failed to backpropagate: %v", err)
	}

	// machine := NewLispMachine(g)  // you can use a LispMachine, but it'll be VERY slow.
	machine = NewTapeMachine(g, BindDualValues(m))
	return m, cost, machine
}

func linregRun(m, cost *Node, machine VM, iter int, autoCleanup bool) []LostAndW {
	if autoCleanup {
		defer machine.Close()
	}
	model := []ValueGrad{m}
	solver := NewVanillaSolver(WithLearnRate(0.01), WithClip(5)) // good idea to clip

	if CUDA {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}
	var err error
	const recordeStripe = config.RecordeStripe4Scalar
	records := make([]LostAndW, 0, iter/recordeStripe)
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Printf("Error during iteration: %v: %v\n", i, err)
			break
		}

		if err = solver.Step(model); err != nil {
			log.Fatal(err)
		}

		if (i+1)%recordeStripe == 0 {
			records = append(records, LostAndW{cost.Value().Data().(float32), m.Value().Data().(float32), 0})
		}
		if (i + 1) == iter {
			fmt.Printf("m: %v  iter: %v Cost: %2.3f  \n",
				m.Value(),
				i+1,
				cost.Value())

			fmt.Println(records)
		}

		machine.Reset() // Reset is necessary in a loop like this
	}
	return records

}

func Train2LearnScalar(sample sample.XY) []LostAndW {
	defer runtime.GC()
	m, cost, machine := linregSetup(sample)
	iter := config.IterScalar
	return linregRun(m, cost, machine, iter, true)
}
