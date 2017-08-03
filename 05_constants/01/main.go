package main

import "fmt"

const someString string = "Gauri"

const (
  oneMoreString = "Go Lang"
  valueOfPI = 3.14
)

func main () {
  // const cannot have short hand declaration const anotherString := "shank" not valid
  const anotherString = "Shankar"

  fmt.Printf(" The value of someString Constant is %v \n",someString)
  fmt.Printf(" The value of anotherString Constant is %v \n",anotherString)
  fmt.Printf(" The value of oneMoreString Constant is %v \n",oneMoreString)
  fmt.Printf(" The value of valueOfPI Constant is %v \n",valueOfPI)
}
