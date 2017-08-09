package main

import "fmt"


// It has to be in a separate go routine
func main() {
  c := make(chan int)
  go func() {
    c <- 10
  } ()

  fmt.Println(<-c)
}
