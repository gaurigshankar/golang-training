package main

import "fmt"

func main() {

  greetings := func(something string) {
    fmt.Println("Hello There",something)
  }

  greetings("Gauri")
  fmt.Printf("%T \n",greetings)
}
