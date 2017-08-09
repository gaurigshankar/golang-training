package main


import "fmt"

func main() {
  numbers := gen(2,3)
  squares := square(numbers)

  for n:= range squares {
    fmt.Println(n)
  }

}

func gen(nums... int) chan int {
  cOut := make(chan int)
  go func(){
    for _, i := range nums {
      cOut <- i
    }

  close(cOut)
  }()
  return cOut
}

func square(c chan int) chan int {
  cOut := make(chan int)
  go func() {
    for n := range c {
      cOut <- n * n
    }
    close(cOut)
  }()
  return cOut
}
