package main

import "fmt"

func main() {
  incrementOutPut := incrementor()
  pullerOut := puller(incrementOutPut)

  for n := range pullerOut {
    fmt.Println(n)
  }
}

func incrementor() chan int{
  cOut := make(chan int)
  go func() {
    for i :=0; i<10; i++ {
      cOut <- i
    }
    close(cOut)
  }()
  return cOut

}


func puller(c chan int) chan int {
  cOut := make(chan int)
  go func() {
    var sum int
    for n := range c {
      sum += n
    }
    cOut <- sum
    close(cOut)
  }()
  return cOut
}
