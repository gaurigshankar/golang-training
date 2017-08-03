package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
  )

// This function is example of blank identifier.
// Since all the variables declared in Go needs tobe used (if not used go wont run the file),
// there would be some cases where we really dont need to use the variable
// In the below example http.Get returns both Response and error
//My use case does not need to use error. In that case if i had used variable to have the error as in the commented lines
// the program may not run. So there should be some alternatives which is where
// blank identifier comes into picture. The blankidentifier is _
// You cannot print blank identifier 


func main () {

  fmt.Println("Printing something ")

  // res, errorFetching := http.Get("http://www.google.com")
  // page, errorWhileRead := ioutil.ReadAll(res.Body)

  res, _ := http.Get("http://www.google.com")
  page, _ := ioutil.ReadAll(res.Body)

  res.Body.Close()
  fmt.Printf("%s", page)

}
