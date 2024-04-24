package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (ver *Vertex) abs() float64 {
  //made abs pointer reciever to avoid copying struct
	return math.Sqrt((ver.X * ver.X) + (ver.Y * ver.Y))
}

func (v *Vertex) change() {
  v.X += 10
  v.Y += 10
}

func main() {
	vertex := Vertex{3, 4}

	fmt.Println(vertex.abs())
  // outupt will be 5

  vertex.change()
  // here change is is a pointer reciever so it can change the value of the reciever itself

	fmt.Println(vertex.abs())
  // outupt changes to 19.something snice the reciever value changed

}
