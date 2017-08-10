package main

import (
  "fmt"
  "net/http"

)

func main() {
  http.HandleFunc("/", helloHandler)
  http.ListenAndServe("localhost:9999", nil)
}


func helloHandler(response http.ResponseWriter, request *http.Request ) {
  fmt.Println("Inside Handler...")
  fmt.Fprintf(response, "Hello from My Go Program")
}
