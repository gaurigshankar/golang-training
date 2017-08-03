package main

import "fmt"
//

var numberValue = 24

func main() {
  fmt.Printf("Value printed from main is %v   \n", numberValue)
  fmt.Printf("numberValueBelow printed from main is %v   \n", numberValueBelow)
  anotherFileFunction()
  foo()
}

func foo() {
  fmt.Printf("Value printed from foo is %v \n", numberValue)
}

var numberValueBelow = 23
