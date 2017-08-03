package main

import "fmt"

func main() {
  b:=4
  if food:="Idly"; len(food) == b {
    fmt.Println("The length is 4 and value of food is ",food)
  }
  fmt.Println("If i try to pring the value of food here, its going to be error out, since, food was defined only in the if block scope",food)
}
