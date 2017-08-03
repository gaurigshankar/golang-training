package main

import "fmt"

func visit(numbers []int, callbackFunction func(int)){
  for _, n := range numbers {
    callbackFunction(n)
  }
}

func main() {
  visit([]int {1,2,3,4}, func(singleNumber int){
    fmt.Println(singleNumber)
  })
}
