package main

import (
  "errors"
  "log"
)

var ErrGauriCustomError = errors.New("gauri first custom error message : trying to find squareroot of negative number")

func main() {
  _, err := mySquareRoot(-10)
  if err != nil {
    log.Fatalln(err)
  }
}


func mySquareRoot(f float64) (float64, error) {
  if f < 0 {
    return 0, ErrGauriCustomError
  }
  // Actual implementation
  return 42, nil
}
