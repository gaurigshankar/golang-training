package main

import "fmt"


func main() {
  fmt.Println(average(1,2,3,4,5))
}


// In this function we have defined marks to be a variadic, meaning this variable can have any number
// of values. Its somewhat close to list / array or commonly collecton
func average(marks ...float64) float64 {

  fmt.Println(marks)
  fmt.Printf("%T \n", marks)

  var total float64

  for _ , mark := range marks {
    total += mark
  }
  return total / float64(len(marks))
}
