package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

// execute the program to see race condition, by
// go run --race main.go

// mutex mutually exclusive / lock
// Before updating the common/shared variable, lock so that only one thread has access to the shared object
// once updating of the common/shared variable is done, Unlock it

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
    mutex.Lock()
    counter ++
    fmt.Println(s, i, "counter : ",counter)
    mutex.Unlock()
  }
  wg.Done()
}
