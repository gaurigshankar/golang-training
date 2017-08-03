package main

import "fmt"

// We can have multiple case match in a single line as here
// the same line gets printed for walmart or amazon

func main() {
  switch "walmart" {
  case "Gauri", "Shankar":
    fmt.Println("Hello Gauri or Shankar")
  case "sanbruno", "sanjose":
    fmt.Println("Hello san bruno or sanjose")
  case "amazon", "walmart":
    fmt.Println("Hello retailers")
  case "Europe", "Africa":
    fmt.Println("Hello continents")
 default:
    fmt.Println("Hello default")
  }


}
