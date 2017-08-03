package main

import "fmt"


func main() {
  nestedLoop()
}


func singleForLoop() {

  for i := 0; i<10; i++ {
    fmt.Println(i)
  }

}

func nestedLoop() {
  for i := 0; i < 5; i++ {
    for j := 0; j < 3; j++ {
      fmt.Println(i," - ",j)
    }
  }
}
