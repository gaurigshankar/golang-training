package main

import "fmt"

var gauri  = "10";


func main() {
  a := 10;
  b := "golang";
  c := 4.17
  d := true

// To print the values uses %v verbs
  fmt.Printf("%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n", d)

  fmt.Printf("%v \n", gauri)



  // To print the type of the variable use %T verbs

  fmt.Printf("%T \n", a)
  fmt.Printf("%T \n", b)
  fmt.Printf("%T \n", c)
  fmt.Printf("%T \n", d)
  fmt.Printf("%T \n", gauri)
}
