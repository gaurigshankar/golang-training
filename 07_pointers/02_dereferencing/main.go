package main

import "fmt"



// to see the content of a pointer varable drefernce pointer variable by prefixing a * infront of it
// Here to view the variable's value b, I Am using *b
func main() {
  a := 44

  fmt.Printf("The value of variable a is %v \n ", a)
  fmt.Printf("The mem address of variable a is %v \n ", &a)

  var b *int = &a;
  fmt.Printf("The value of variable b is %v \n ", b)
  fmt.Printf("The dereferenced value of variable b is %v \n ", *b)
}
