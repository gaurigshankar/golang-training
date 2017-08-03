package main

import "fmt"
import "github.com/gaurigshankar/golang-training/01_helloworld/stringutils"

func main() {
	fmt.Println("Hello World")
  fmt.Println("Hello "+stringutils.MyName)
  fmt.Println(stringutils.Reverse("Hello "))
}
