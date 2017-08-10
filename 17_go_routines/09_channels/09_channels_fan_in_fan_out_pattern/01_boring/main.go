package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  c := fanIn(boring("Gauri"),boring("shankar"))
  for i := 0 ; i< 100; i++ {
    fmt.Println(<-c)
  }
  fmt.Println("You both are Boring and I am leaving")
}

func boring(msg string) <- chan string{
  cOut := make(chan string)
  go func() {
    for i :=0 ;; i++ {
      cOut <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()
  return cOut
}

func fanIn(input1, input2 <- chan string) <- chan string {
  cOut := make(chan string)

  go func() {
    for {
      cOut <- <-input1
    }
  }()

  go func() {
    for {
    cOut <- <-input2
  }
  }()

  return cOut
}
