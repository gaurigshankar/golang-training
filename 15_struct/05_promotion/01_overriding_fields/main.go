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
  fName string
}

func main() {
  e1 := Employee {
    Person: Person{
      fName: "Gauri",
      lName: "Shankar",
      age: 30,
    },
    employementStatus : true,
    fName: "Shankar",

  }

  e2 := Employee {
    Person: Person{
      fName: "Ram",
      lName: "Kumar",
      age: 40,
    },
    employementStatus : false,
    fName: "Kumar",
  }

// overiding values
  fmt.Println("Child field ",e1.fName, " parent field ", e1.Person.fName)
  fmt.Println("Child field ",e2.fName, " parent field ",e2.Person.fName)
}
