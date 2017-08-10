package main

import (
    "golang.org/x/net/websocket"
    "net/http"
    "fmt"
)

func echoHandler(ws *websocket.Conn) {

  for {
    receivedtext := make([]byte, 100)

    n,err := ws.Read(receivedtext)

    if err != nil {
      fmt.Printf("Received: %d bytes\n",n)
    }

    s := string(receivedtext[:n])
    fmt.Printf("Received: %d bytes: %s\n",n,s)
    //io.Copy(ws, ws)
    //fmt.Printf("Sent: %s\n",s)
  }
}

func main() {
  http.Handle("/echo", websocket.Handler(echoHandler))
  http.HandleFunc("/d", handleChatPage)
  err := http.ListenAndServe(":9999", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}

func handleChatPage(response http.ResponseWriter, request *http.Request) {
  http.ServeFile(response, request, "../07_webscoket/websocket-chat.html")
}
