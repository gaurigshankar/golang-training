package main

import "fmt"


func main() {
  numbersGenerated := gen()
  factorialOutput := calculateFactorial(numbersGenerated)

  for n := range factorialOutput {
    fmt.Println(n)
  }
}

func gen() chan int{
  cOut := make(chan int)
  go func() {
      for i := 0; i <10; i++ {
        for j:=3; j<13; j++ {
          cOut <- j
        }
      }
      close(cOut)
  }()
  return cOut
}

func calculateFactorial(input chan int) chan int {
  cOut := make(chan int)
  go func () {
      for n := range input{
        cOut <- factorial(n)
      }
      close(cOut)
  }()
  return cOut
}

func factorial(input int) int {

    product := 1
    for j := input; j>0; j-- {
      product *= j
    }
    
  return product
}
