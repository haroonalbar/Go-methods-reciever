package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (ver Vertex) abs() float64 {
	return math.Sqrt((ver.X * ver.X) + (ver.Y * ver.Y))
}

func main() {
	vertex := Vertex{3, 4}

	fmt.Println(vertex.abs())

}
