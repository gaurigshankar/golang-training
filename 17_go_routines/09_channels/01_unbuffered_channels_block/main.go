package main

import (
  "fmt"
  "time"
)


//channel works like relay race. Where the person1 needs to handle over the baton to another

func main () {
  myDummyChannel := make(chan int)

  go func() {
    for i :=0; i<20; i++ {
      myDummyChannel <- i
    }
  }()
  go func() {
    for {
      fmt.Println(<-myDummyChannel)
    }
  }()
  time.Sleep(time.Second)
}
