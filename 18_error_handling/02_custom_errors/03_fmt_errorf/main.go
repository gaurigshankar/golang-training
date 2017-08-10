package main

import (
  "fmt"
  "log"
)

// instead of error.New using fmt.Errorf so we can provide context
// in error.New we can just pass over the string
// in fmt.Errorf we can append some dynamic variables to the error message
// so error message is more valuable for debugging

func main() {
  _, err := mySquareRoot(-10)
  if err != nil {
    log.Fatalln(err)
  }
}


func mySquareRoot(f float64) (float64, error) {
  if f < 0 {
    return 0, fmt.Errorf("gauri first custom error message : trying to find squareroot of negative number: %v",f)
  }
  // Actual implementation
  return 42, nil
}
