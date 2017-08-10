package main

import (
  "os"

)


//Package log implements a simple logging package..... writes to standard error and prints the
//date and time of each logged message
// the fatal functions call os.Exit(1) after writing the log messange
// the panic functions call panic after writing the log message

// panic basically gives more stack information
func main() {
  panicError()

}

func panicError(){
  _, err := os.Open("non-existing_file.txt")
  if err != nil {
    panic("Err happened and logged using log" )
  }
}
