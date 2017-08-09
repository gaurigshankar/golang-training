package main

import (
  "encoding/json"
  "fmt"
)


// This Struct has tags for fields
// LName filed need not be exposed outside so tagging with a -
// Age fied . This field name to be renamed as namingAgeFieldAsSomethingElse when marshaled 
type Person struct {
  FName string
  LName string `json:"-"`
  Age int `json:"namingAgeFieldAsSomethingElse"`
}


func main() {
  p1 := Person{"Gauri", "shankar",  30}
  //p2 := AnothPersonWithNonExportedFields{"Ram", "Kumar",  40}

  byteStreamForPerson, _ := json.Marshal(p1);
  //byteStreamForAnotherPerson, _ := json.Marshal(p2);

  fmt.Println(string(byteStreamForPerson));
  // Since the AnothPersonWithNonExportedFields has non-exported fields marshal will not give output
  //fmt.Println(string(byteStreamForAnotherPerson));
}
