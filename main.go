package main

import (
	"fmt"
	"math"
)

type Abster interface{
  abs() float64
}
// interface can be used to handle dynamic functions
// here there are 2 types of abs() one that takes MyFloat and other that takes *Vertex
// and Abster interface handles both abs methods.

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) abs() float64 {
  if(f < 0){
    return -float64(f)
  }else{
    return float64(f)
  }
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
  var a Abster
  f := MyFloat(-0.111)
  a = f
	fmt.Println(a.abs())

	vertex := Vertex{3, 4}
  a = &vertex

	fmt.Println(a.abs())
  // outupt will be 5

  vertex.change()
  // here change is is a pointer reciever so it can change the value of the reciever itself

	fmt.Println(a.abs())
  // outupt changes to 19.something snice the reciever value changed

}
