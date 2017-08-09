package main

import (
  "encoding/json"
  "fmt"
)

// All the fields starts with small letter meaning not exported
type AnothPersonWithNonExportedFields struct {
  fName string
  lName string
  age int
}

type Person struct {
  FName string
  LName string
  Age int
}


func main() {
  p1 := Person{"Gauri", "shankar",  30}
  p2 := AnothPersonWithNonExportedFields{"Ram", "Kumar",  40}

  byteStreamForPerson, _ := json.Marshal(p1);
  byteStreamForAnotherPerson, _ := json.Marshal(p2);

  fmt.Println(string(byteStreamForPerson));
  // Since the AnothPersonWithNonExportedFields has non-exported fields marshal will not give output
  fmt.Println(string(byteStreamForAnotherPerson));
}
