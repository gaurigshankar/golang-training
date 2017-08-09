package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
  "sync/atomic"
)

var counter int64
var wg sync.WaitGroup
var mutex sync.Mutex

// execute the program to see race condition, by
// go run --race main.go

// Instead of locking we are using AtomicInteger //AtomicVariables in Java
// we are using int64 as the type for counter and atomic.AddInt64 for adding
// so that the whole section will not be locked


func main() {
  wg.Add(2)
  go incrementor("foo")
  go incrementor("bar")
  wg.Wait()
  fmt.Println("Final Counter ",counter)
}

func incrementor (s string) {
  for i:=0; i<25; i++ {
    time.Sleep(time.Duration(rand.Intn(3))* time.Millisecond)
    atomic.AddInt64(&counter, 1)
    fmt.Println(s, i, "counter : ",counter)
  }
  wg.Done()
}
