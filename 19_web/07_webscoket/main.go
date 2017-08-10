package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/websocket"
  "github.com/gorilla/mux"
)

var muxRouter = mux.NewRouter()
var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func main() {
  muxRouter.HandleFunc("/echo", handleEcho)
  muxRouter.HandleFunc("/", handleChatPage)
  http.ListenAndServe(":9999", muxRouter)
}

func handleEcho(response http.ResponseWriter, request *http.Request) {
  conn, _ := upgrader.Upgrade(response, request, nil)

  for {
    msgType, msg, err := conn.ReadMessage()

    if err != nil {
      return
    }

    fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

    if err = conn.WriteMessage(msgType, msg); err != nil {
      return
    }
  }

}

func handleChatPage(response http.ResponseWriter, request *http.Request) {
  http.ServeFile(response, request, "websocket-chat.html")
}
