package main

import "fmt"
import "github.com/gaurigshankar/golang-training/03_scopes/02_visibility/vis"

func main() {
  fmt.Printf("Hi hello \n")
  fmt.Printf(" The name refered from different package is %v \n", vis.MyName)
  //fmt.Printf(" The title refered from different package is %v and may not be printed due to non-export \n", vis.myTitle)
  vis.GauriCustomPrinterFunction()
}
