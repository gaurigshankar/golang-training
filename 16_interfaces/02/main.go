package main

import "fmt"
import "math"


type Square struct {
  side float64
}

type Circle struct {
  radius float64
}

type Shape interface {
  area() float64
}

func (s Square) area() float64 {
  return s.side * s.side
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func info(shape Shape) {
  fmt.Println("which shape? ",shape)
  fmt.Println("area of the shape ", shape.area())

}

func main() {
  square := Square{10}
  circle := Circle{20}
  info(square)
  info(circle)
}
