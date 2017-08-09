package main

import "fmt"

type person struct {
  fName string
  lName string
  age int
}

func (p person)  fullName() string {
  return p.fName + p.lName
}

func main() {
  p1 := person{"Gauri", "shankar",  30}
  p2 := person{"Ram", "Kumar",  40}

  fmt.Println(p1.fullName())
  fmt.Println(p2.fullName())
}
