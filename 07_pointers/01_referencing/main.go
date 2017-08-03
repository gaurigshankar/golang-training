package main

import "fmt"



// This program is demo showing referncing a variable to a memory address of another variable
// As usual the memory address of a vaiable can be got prefixin & infront of the variable
// Incase if you want to refer the pointer of the variable to another variable you could declare the variable
// to be pointer of given type as like var variableName *typeOfVariable. Note there is a * symbol infront of the typeOfVariable
// This says this variable is a pointer not acutal value
// So in the below example we have declared a pointer variable of type int called b and assigned the meemory address of a
func main() {
  a := 43

  fmt.Printf("The value of variable a is %v \n ", a)
  fmt.Printf("The mem address of variable a is %v \n ", &a)

  var b *int = &a;
  fmt.Printf("The value of variable b is %v \n ", b)
}
