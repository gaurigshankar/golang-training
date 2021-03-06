package main

import "fmt"
import "sync"

var wg sync.WaitGroup

//nothing will be printed to the console because main exists
func main() {
  fmt.Println("Starting main and spawns separate thread for foo and bar")
  wg.Add(2)
  go foo()
  go bar()
  wg.Wait()
  fmt.Println("Main thread exists")
}

func foo() {
  for i := 0 ; i < 45; i ++ {
    fmt.Println("Foo: ",i)
  }
  wg.Done()
}

func bar() {
  for i := 0 ; i < 45; i ++ {
    fmt.Println("Bar: ",i)
  }
  wg.Done()
}
