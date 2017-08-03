package main

import "fmt"


func main() {
 output :=  filter([]int {1,2,3,4,5}, func(someNumber int) bool {
    return someNumber > 1
  })
  fmt.Println(output)
}


func filter (numbers []int, callbackFunction func(n int) bool) [] int {
  filterdNumbers := []int{}

  for _, singleNumber := range numbers {
    if callbackFunction(singleNumber) {
      filterdNumbers = append(filterdNumbers, singleNumber)
    }
  }
  return filterdNumbers
}
