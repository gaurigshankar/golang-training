package main

import (
  "os"
  "fmt"
  "log"
)


//Package log implements a simple logging package..... writes to standard error and prints the
//date and time of each logged message
// the fatal functions call os.Exit(1) after writing the log messange
// the panic functions call panic after writing the log message


func main() {


  logStandardError()
}



func logStandardError() {
  _, err := os.Open("non-existing_file.txt")
  if err != nil {
    log.Println("Err happened and logged using log", err )
  }
  fmt.Println("Continue working after logging ......")
}
