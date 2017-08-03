package main

import "fmt"


func main() {
  switch "Gauri" {
  case "Shankar" :
    fmt.Println("Hello Shankar")
  case "amazon" :
    fmt.Println("Hello amazon")
  case "sanbruno" :
    fmt.Println("Hello san bruno")
  case "Gauri" :
    fmt.Println("Hello gauri")
  default:
    fmt.Println("Hello Default")
  }

}
