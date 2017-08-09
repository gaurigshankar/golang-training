package main

import "fmt"


func main() {
  yetAnotherMap := map[string] string {"fName" : "gauri", "lName" : "shankar" , "place" :
  "Antartica"}

  for key, val := range yetAnotherMap {
    fmt.Println(" key ",key, " val ",val)
  }
}
