package main

import "fmt"

func main() {
  fmt.Println(greet("gauri", "shankar"))
}

func greet(fname, lname string) (concatenatedString string) {
  //fmt.Sprint is concatenates given strings
  // Here when we declared the return type we have defined the return variable name in func declaration
  // So just doing an empty return returns the variable defined
  concatenatedString = fmt.Sprint(fname," ",lname)
  return
}
