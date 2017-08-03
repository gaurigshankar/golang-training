package main

import "fmt"


// in main function world() is prepended with defer keyword
// defer keyword defers the execution of the line, till all the statements in the current function executes

func main() {
  defer world()
  hello()
}

func hello () {
  fmt.Print("Hello ")
}

func world () {
  fmt.Println("World")
}
