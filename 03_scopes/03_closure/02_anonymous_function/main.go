package main

import "fmt"




func main() {
  x := 0;
  fmt.Printf("The initial value of X is %v \n",x)
  increment := func () int {
    x++
    return x
  }

  fmt.Println(increment())
  fmt.Println(increment())

}
