package main

import "fmt"

type person struct {
  fName string
  lName string
  age int
}

func main() {

  p1 := person{"Gauri", "shankar",  30}
  p2 := person{"Ram", "Kumar",  40}

  fmt.Println("p1 ",p1)
  fmt.Println("p2 ",p2)


  anotherTypeInitialization()
}

// Another Type. Note the there is an extra comma after age value, that is absolutely needed
// else program will not run

func anotherTypeInitialization() {
  p1 := person{
    fName: "Gauri",
    lName: "shankar",
    age: 30,
  }
  p2 := person{
    fName: "Ram",
    lName: "Kumar",
    age: 40,
  }

  fmt.Println("p1 ",p1)
  fmt.Println("p2 ",p2)
}
