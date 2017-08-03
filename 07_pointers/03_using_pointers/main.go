package main

import "fmt"

// We can reassign the value stored in a memory address as *b = 40.
// Here the memory Address value gets re assigned to 40
// The particiular memory address is being referenced by variable a
// So as expected the new value of a will be 40. this is proven by this example

func main() {
  a := 44

  fmt.Printf("The value of variable a is %v \n ", a)
  fmt.Printf("The mem address of variable a is %v \n ", &a)

  var b *int = &a;
  fmt.Printf("The value of variable b is %v \n ", b)
  fmt.Printf("The dereferenced value of variable b is %v \n ", *b)

  *b = 40

  fmt.Printf("The value of variable a after modifying the value stored in the addresss pointed by b is %v \n ", a)
}
