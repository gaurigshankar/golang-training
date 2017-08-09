package main

import "fmt"

func main() {
  incrementOutPutFoo := incrementor("Foo :")
  incrementOutPutBar := incrementor("Bar :")
  pullerOutFoo := puller(incrementOutPutFoo)
  pullerOutBar := puller(incrementOutPutBar)

  fmt.Println("The final output ", <-pullerOutFoo+<-pullerOutBar)

}

func incrementor(s string) chan int{
  cOut := make(chan int)
  go func() {
    for i :=0; i<10; i++ {
      cOut <- i
      fmt.Println(s, i)
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
