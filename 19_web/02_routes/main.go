package main

import (
  "fmt"
  "net/http"
)

var userages = map[string] int {
  "Gauri" : 30,
  "Shankar": 40,
  "Ram": 35,
  "Kumar": 45,
}

func main() {
  http.HandleFunc("/users/", handleUsersAge)
  http.ListenAndServe("localhost:9999", nil)

}

func handleUsersAge(response http.ResponseWriter, request *http.Request){


  name := request.URL.Path[len("/users/"):]
  age := userages[name]
  fmt.Fprintf(response, "%s is %d years old", name, age)
}
