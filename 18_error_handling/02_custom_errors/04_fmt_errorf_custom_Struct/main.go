package main

import (
  "fmt"
  "log"
)


type GauriMathError struct {
  level string
  isSevere bool
  err error
}

func (newError GauriMathError) Error() string {
  return fmt.Sprintf("Application Error occured with level %v and severity of %v and the error message is %v",newError.level, newError.isSevere, newError.err)
}


func main() {
  _, err := mySquareRoot(-10)
  if err != nil {
    log.Fatalln(err)
  }
}


func mySquareRoot(f float64) (float64, error) {
  if f < 0 {
    someErrorMessage := fmt.Errorf("gauri first custom error message : trying to find squareroot of negative number: %v",f)
    return 0, &GauriMathError{"info", true, someErrorMessage}
  }

  // Actual implementation
  return 42, nil
}
