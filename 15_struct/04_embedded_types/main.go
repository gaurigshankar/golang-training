package main

import "fmt"

type Person struct {
  fName string
  lName string
  age int
}

func (p Person)  fullName() string {
  return p.fName + p.lName
}

type Employee struct {
  Person
  employementStatus bool
}

func main() {
  e1 := Employee {
    Person: Person{
      fName: "Gauri",
      lName: "Shankar",
      age: 30,
    },
    employementStatus : true,
  }

  e2 := Employee {
    Person: Person{
      fName: "Ram",
      lName: "Kumar",
      age: 40,
    },
    employementStatus : false,
  }

  fmt.Println(e1)
  fmt.Println(e2)
  fmt.Println(e1.fullName())
  fmt.Println(e2.fullName())
}
