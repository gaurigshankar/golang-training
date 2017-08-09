package main

import "fmt"


func main() {
  factorialValue := factorial(4)

  for n := range(factorialValue) {
    fmt.Println(n)
  }

}

func factorial(n int) chan int{
  cOut := make(chan int)

  go func() {
    var product int
    product = 1
    for i := n; i>0; i-- {
      product *= i
    }
    cOut <- product
    close(cOut)
  }()

  return cOut
}
