package main

import "fmt"

// There is no need to give the expression to the switch
// You can evaluate them inside the case

func main() {

  myFriendsName := "Ramu"

  switch  {
  case len(myFriendsName) == 2:
    fmt.Println("Hello myFriendsName has length of 2")
  case myFriendsName == "Tim":
    fmt.Println("Hello My friend name is Tim")
  case myFriendsName == "Jenny":
    fmt.Println("Hello My friend name is Jenny")
  case myFriendsName == "Something":
    fmt.Println("Hello My friend name is  Something")
 default:
    fmt.Println("Hello default")
  }


}
