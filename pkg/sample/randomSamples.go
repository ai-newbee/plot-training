package sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	c "plot-training/pkg/config"
)

type XY struct {
	X []float32 `json:"x"`
	Y []float32 `json:"y"`
	B []float32 `json:"base"`
}

func New(count int, slope, base float32) XY {
	var result = XY{}
	result.X = make([]float32, 0, count)
	result.Y = make([]float32, 0, count)

	seed := rand.NewSource(314 * 33 * 21) //
	r := rand.New(seed)

	for i := 0; i < count; i++ {
		x := r.Float32()
		result.X = append(result.X, x)

		smallError := r.Float32() / 10
		var y float32
		if base == 0 {
			y = (slope + smallError) * x
		} else {
			y = (slope+smallError)*x + (base + smallError)
			result.B = append(result.B, base+smallError)
		}
		//y := slop * x
		result.Y = append(result.Y, y)
	}
	//save2file(result)
	return result
}

func save2file(samples XY) {
	bytes, _ := json.MarshalIndent(samples, "", " ")
	//json := string(bytes)
	//fmt.Printf("%v \n", json)

	err := ioutil.WriteFile(c.SampleFilePath, bytes, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Printf("filePath name is %s \n", c.SampleFilePath)
}
