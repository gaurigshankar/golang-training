package main

import "fmt"

func main() {
  var a [64]string
  fmt.Println(len(a))
  fmt.Println(a)

  for i :=65; i < 128; i++ {
    a[i-65 ] = string(i)
  }
  fmt.Println(a)

}
