package main

import "fmt"


//Determine why this results in deadlock
func main() {
  c := make(chan int)
  c <- 10
  fmt.Println(<-c)
}
