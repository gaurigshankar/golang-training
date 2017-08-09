package main

import "fmt"

// why does it print only zero 0
// what needs to be changed to print from 0 to 9


func main() {
  c := make(chan int)
  go func() {
    for i := 0; i <10; i++ {
      c <- i
    }
  }()

    fmt.Println(<-c)

}
