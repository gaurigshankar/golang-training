package main

import "fmt"

func main() {
  dummySlice :=  make([]string, 3, 8)

  fmt.Println("dummySlice ",dummySlice)
  fmt.Println("dummySlice length ",len(dummySlice))
  fmt.Println("dummySlice  capacity", cap(dummySlice))

  dummySlice[0] = "gauri"
  dummySlice[1] = "shankar"
  dummySlice[2] = "amazon"
  //need to use append when the size exxccedds
  dummySlice = append(dummySlice, "sanbruno")
  fmt.Println("dummySlice ",dummySlice)
  fmt.Println("dummySlice length ",len(dummySlice))
  fmt.Println("dummySlice  capacity", cap(dummySlice))
  fmt.Println("dummySlice  at 4", dummySlice[3])

  fmt.Println("dummySlice  at 4", dummySlice[1:3]) // slicing a slice

  twoDimensionalArray()
}


func twoDimensionalArray() {
  transactions := make([][]int, 5)
  for i := 0 ; i < 4; i ++ {
    transaction := make([]int, 5)
    for j:=0; j< 3; j++ {
      transaction[j] = j;
    }
    transactions[i] = transaction
  }

  fmt.Println(transactions)
}
