package main

import "fmt"

func anotherFileFunction() {
  fmt.Printf("Value printed from anotherFileFunction is %v \n", numberValue)
  foo()
  square := calculateSquare(7)
  fmt.Printf("Squared Output is %v \n", square)
}

func calculateSquare(x int) int {
  return x * x
}
