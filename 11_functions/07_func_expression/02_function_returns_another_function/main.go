package main

import "fmt"


func main() {
  greet := makeGreeter()
  fmt.Println(greet())
}


// Here this function returns another function that inturn retuns string
func makeGreeter() func() string {
  return func() string {
    return "Hello World"
  }
}
