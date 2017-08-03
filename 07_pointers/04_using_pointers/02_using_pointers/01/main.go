package main

import "fmt"

func main() {

  x := 5
  fmt.Printf(" The value of x before calling zero function is %v \n", x)
  fmt.Printf(" The mem address of x before calling zero function is %v \n", &x)
  zero(&x)

  fmt.Printf(" The value of x after calling zero function is %v \n", x)
  fmt.Printf(" The mem address of x after calling zero function is %v \n", &x)


}


func zero(z *int) {
  *z = 0
}
