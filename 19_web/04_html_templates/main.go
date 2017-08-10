package main

import (
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

type ToDo struct {
  Task string
  Done bool
}

var muxRouter = mux.NewRouter()

func main() {



  muxRouter.HandleFunc("/", handleToDo).Methods("GET")
  http.ListenAndServe(":9999", muxRouter)
}

func handleToDo(response http.ResponseWriter, request *http.Request) {

  todos := []ToDo{
    {"Learn Go", true},
    {"Read Go Examples", true},
    {"Create Webapp in GO", false},
  }

  tmpl := template.Must(template.ParseFiles("todos.html"))
  tmpl.Execute(response, struct{Todos []ToDo}{todos})

}
