package sample

import (
	c "dl-base/pkg/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type XY struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func Random(count int) []XY {
	var result = make([]XY, 0, count)
	seed := rand.NewSource(314 * 33 * 21) //
	r := rand.New(seed)
	for i := 0; i < count; i++ {
		x := r.Float32()
		y := r.Float32()
		xy := XY{x, y}
		result = append(result, xy)
	}
	save2file(result)
	return result
}

func save2file(samples []XY) {
	bytes, _ := json.MarshalIndent(samples, "", " ")
	//json := string(bytes)
	//fmt.Printf("%v \n", json)

	err := ioutil.WriteFile(c.SampleFilePath, bytes, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Printf("filePath name is %s", c.SampleFilePath)
}
