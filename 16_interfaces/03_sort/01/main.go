package main

import (
  "fmt"
  "sort"
)


type people [] string

func (p people) Len() { return len(p) }
func (p people) Swap(i, j int) { return p[i] p[j] = p[j] p[i]}
func (p people) Less(i, j int) bool  { return p[i] < p[j] }


func main() {
  p := people{"zen", "la", "gauri", "al"}
  s := []string {"zen", "la", "gauri", "al"}
  i := []int {40, 11, 1, 96, 54}
  sort.Sort(p)
  fmt.Println(p)


  sort.Strings(s)
  fmt.Println(s)

  sort.Sort(sort.IntSlice(i))
  fmt.Println("Ascending order ",i)

  sort.Sort(sort.Reverse(sort.IntSlice(i)))
  fmt.Println("Descending order ",i)

  sort.Ints(i)
  fmt.Println("Ascending order ",i)

}
