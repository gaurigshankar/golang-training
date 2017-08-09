package main

import "fmt"

type Person struct {
  fName string
  lName string
  age int
}

func (p Person)  fullName() string {
  return " from Parent(Person) "+ p.fName + p.lName
}

func (e Employee)  fullName() string {
  return " from child(Employee) " +e.fName + e.lName
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
  fmt.Println("Child field ",e1.fullName(), " parent field ", e1.Person.fullName())
  fmt.Println("Child field ",e2.fullName(), " parent field ",e2.Person.fullName())
}
