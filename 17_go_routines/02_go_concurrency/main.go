package main

import "fmt"

//nothing will be printed to the console because main exists
func main() {
  fmt.Println("Starting main and spawns separate thread for foo and bar")
  go foo()
  go bar()
  fmt.Println("Main thread exists")
}

func foo() {
  for i := 0 ; i < 45; i ++ {
    fmt.Println("Foo: ",i)
  }
}

func bar() {
  for i := 0 ; i < 45; i ++ {
    fmt.Println("Bar: ",i)
  }
}
