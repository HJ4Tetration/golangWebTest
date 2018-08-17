package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type myServer struct {
}

type message struct {
	data1 int
	data2 string
}

type rmessage struct {
	data1 string
	data2 bool
}
type sv struct {
	upgrader websocket.Upgrader
}

/*func (s sv) ServeHTTP(w http.ResponseWriter, request *http.Request) {
}*/

/*func (s *sv) wsHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, request *http.Request) {
		})
}*/

func (s *sv) wsHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("received websocket request from client\n")
	switchConnection, err := s.upgrader.Upgrade(w, request, nil)
	if err != nil {
		fmt.Printf("cannot upgrade the websocket\n")
		return
	}
	go func() {
		for {
			messageType, mess, err := switchConnection.ReadMessage()
			if err != nil {
				fmt.Printf("error reading message from client\n")
				return
			}
			var jsonMessage message
			err = json.Unmarshal(mess, &jsonMessage)
			if err != nil {
				fmt.Printf("cannot unmarshal message from client\n")
				return
			}
			mes2send := rmessage{"hello client", false}
			if jsonMessage.data1 == 5 {
				mes2send.data2 = true
			}
			send, err := json.Marshal(mes2send)
			if err != nil {
				fmt.Printf("cannot marshal message to client\n")
				return
			}
			switchConnection.WriteMessage(messageType, send)
		}
	}()
}

func main() {
	s := sv{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}}
	port := "localhost:9999"
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(s.wsHandler))
	fmt.Printf("trying to start server\n")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Printf("cannot start listenAndServe\n")
		return
	}
}
