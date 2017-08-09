package main

import "fmt"
import "encoding/json"

type Person struct {
  FName string
  LName string
  Age int
}

// when unmarshaling pointer variable needs to provided.
func main() {
  var p1 Person

  bs := []byte(`{"FName": "Gauri", "LName": "Shankar", "Age": 30}`)
  json.Unmarshal(bs, &p1)

  fmt.Println(p1)
}
