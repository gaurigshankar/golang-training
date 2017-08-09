package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)

var counter int
var wg sync.WaitGroup

// execute the program to see race condition, by
// go run --race main.go


// Theoritacally the final value of the counter should be 50, as we are calling
// incrementor twice and each call has 25 iterations
// But you will see less than that as final Counter value due to race condition

// Read more about Race condition in https://github.com/ardanlabs/gotraining/blob/5bb3224966589b99692efff0d8a9492e974e051f/topics/go/concurrency/data_race/README.md

func main() {
  wg.Add(2)
  go incrementor("foo")
  go incrementor("bar")
  wg.Wait()
  fmt.Println("Final Counter ",counter)
}

func incrementor (s string) {
  for i:=0; i<25; i++ {
    x := counter
    x++
    time.Sleep(time.Duration(rand.Intn(3))* time.Millisecond)
    counter = x
    fmt.Println(s, i, "counter : ",counter)
  }
  wg.Done()
}
