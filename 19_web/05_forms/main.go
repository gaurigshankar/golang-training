package main

import (
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

type ContactDetails struct {
  Email string
  Subject string
  Message string
}

var muxRouter = mux.NewRouter()

func main() {
  muxRouter.HandleFunc("/", handleEmailForm)
  http.ListenAndServe(":9999", muxRouter)
}

func handleEmailForm(response http.ResponseWriter, request *http.Request) {
  tmpl := template.Must(template.ParseFiles("forms.html"))
  if(request.Method != http.MethodPost) {
    tmpl.Execute(response, nil)
    return
  }
  details := ContactDetails{
    Email: request.FormValue("email"),
    Subject: request.FormValue("Subject"),
    Message: request.FormValue("Message"),
  }

  _ = details

  tmpl.Execute(response, struct{ Success bool}{true})
}
