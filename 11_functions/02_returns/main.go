package main

import "fmt"

func main() {
  fmt.Println(greet("gauri", "shankar"))
}

func greet(fname, lname string) string {
  //fmt.Sprint is concatenates given strings
  return fmt.Sprint(fname," ",lname)
}
