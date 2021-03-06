package main

import (
  "fmt"
  "sync"
  )

func main() {
  myDummyChannel := make(chan int)

  var wg sync.WaitGroup

  go func() {
      wg.Add(1)
      for i:=0; i <10 ; i++ {
        myDummyChannel <- i
      }
      wg.Done()
  }()

  go func() {
      wg.Add(1)
      for i:=0; i <10 ; i++ {
        myDummyChannel <- i
      }
      wg.Done()
  }()

  go func() {
    wg.Wait()
    close(myDummyChannel)
  }()

  for n:= range myDummyChannel {
    fmt.Println(n)
  }
}
