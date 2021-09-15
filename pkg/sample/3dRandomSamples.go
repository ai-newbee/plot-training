package sample

import "C"
import (
	"encoding/csv"
	"strconv"

	"fmt"
	"math/rand"
	"os"
	c "plot-training/pkg/config"
)

type XYZ struct {
	X []float32 `json:"x"`
	Y []float32 `json:"y"`
	Z []float32 `json:"z"`
}

func New3DSample(count int, csvFileName string) (result XYZ) {
	result = XYZ{}
	result.X = make([]float32, 0, count)
	result.Y = make([]float32, 0, count)
	result.Z = make([]float32, 0, count)

	seed := rand.NewSource(314 * 33 * 21) //
	r := rand.New(seed)

	for i := 0; i < count; i++ {
		x := r.Float32() * 2
		result.X = append(result.X, x)
		y := r.Float32() * 2
		result.Y = append(result.Y, y)

		var z float32
		c := []float32{1.0, 1.0} // center of a cycle
		if (x-c[0])*(x-c[0])+(y-c[1])*(y-c[1]) < 0.2 {
			z = 1
		} else {
			z = 0
		}
		result.Z = append(result.Z, z)
	}
	saveAsCSV(result, csvFileName)
	return result
}

func saveAsCSV(samples XYZ, csvFileName string) {
	absFilePath := c.DatasetDir + csvFileName
	os.Create(absFilePath) //clean the content of csv file
	f, err := os.OpenFile(absFilePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"x", "y", "z"}
	writer.Write(header)

	for i := 0; i < len(samples.X); i++ {
		var data = []string{f2s(samples.X[i]), f2s(samples.Y[i]), f2s(samples.Z[i])}
		writer.Write(data)
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("filePath name is %s \n", absFilePath)
}

func f2s(fv float32) string {
	return strconv.FormatFloat(float64(fv), 'f', 10, 32)
}
