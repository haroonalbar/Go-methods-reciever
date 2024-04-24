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

// func (v *Vertex) change() {
//   v.X += 10
//   v.Y += 10
// }
func  change(v *Vertex) {
  v.X += 10
  v.Y += 10
}

func (v Vertex) nochange() {
  v.X += 10
  v.Y += 10
}

func main() {
	vertex := Vertex{3, 4}

	fmt.Println(vertex.abs())
  // outupt will be 5

  vertex.nochange()
  // since nochage is not a pointer reciever it will not change the value of the reciever
  // so abs would be the same as  the above in this case 5
	fmt.Println(vertex.abs())

  change(&vertex)
  // here change is is a pointer reciever so it can change the value of the reciever itself

	fmt.Println(vertex.abs())
  // outupt changes to 19.something snice the reciever value changed

}
