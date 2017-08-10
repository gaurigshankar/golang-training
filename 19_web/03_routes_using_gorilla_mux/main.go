package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

var muxRouter = mux.NewRouter()
var userages = map[string] int {
  "Gauri" : 30,
  "Shankar": 40,
  "Ram": 35,
  "Kumar": 45,
}

func main() {
  muxRouter.HandleFunc("/users/{name}", handleUsersAge).Methods("GET")
  http.ListenAndServe(":9999", muxRouter)
}


func handleUsersAge(response http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  name := vars["name"]
  age := userages[name]

  fmt.Fprintf(response, "%s is %d years old", name, age)
}
