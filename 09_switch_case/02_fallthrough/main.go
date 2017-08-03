package main

import "fmt"

// fallthrough  means continue till u see a break; Remember break is by default applied to all cases
// In the below example, output will be hello amazon, hello san bruno , hello gauri.
// In the case definiation for gauri, there is no fallthrough specified, so default break will be appplied and
// flow terminates
func main() {
  switch "amazon" {
  case "Shankar" :
    fmt.Println("Hello Shankar")
  case "amazon" :
    fmt.Println("Hello amazon")
    fallthrough
  case "sanbruno" :
    fmt.Println("Hello san bruno")
    fallthrough
  case "Gauri" :
    fmt.Println("Hello gauri")
  default:
    fmt.Println("Hello Default")
  }
}
