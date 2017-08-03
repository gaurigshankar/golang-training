package main

import "fmt"

func main() {
    x := 10

    SwitchOnType(x)
    SwitchOnType("Hello Gauri")
    var newContact = Contact {"Hello", "Gauri"}
    SwitchOnType(newContact)
}

type Contact struct {
  greetings string
  name string
}


func SwitchOnType(x interface{}){
  switch x.(type) {
  case int:
    fmt.Println("Int")
  case string:
    fmt.Println("String")
  case Contact:
    fmt.Println("Contact")
  default:
      fmt.Println("Default")
  }
}
