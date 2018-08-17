package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type myClient struct {
	d  websocket.Dialer
	ip string
}

type message struct {
	data1 int    `json:"data1"`
	data2 string `json:"data2"`
}

type rmessage struct {
	data1 string `json:"data1"`
	data2 bool   `json:"data2"`
}

func main() {
	mc := myClient{ip: "ws://127.0.0.1:9999/"}
	serverConnection, _, err := mc.d.Dial(mc.ip, nil)
	if err != nil {
		fmt.Printf("error making connection to server")
		return
	}
	mes2send := message{0, "hello,server"}
	for i := 1; i < 10; i++ {
		mes2send.data1++
		send, err := json.Marshal(mes2send)
		if err != nil {
			fmt.Printf("cannot marshal message (client to server)")
			return
		}
		err = serverConnection.WriteMessage(2, send)
		if err != nil {
			fmt.Printf("cannot send message to server")
			return
		}
		fmt.Printf("message sent %v\n", mes2send)
		_, receive, err := serverConnection.ReadMessage()
		if err != nil {
			fmt.Printf("cannot receive message from server")
			return
		}
		var mes2rec rmessage
		err = json.Unmarshal(receive, &mes2rec)
		if err != nil {
			fmt.Printf("cannot Unmarshal message from server")
			return
		}
		if mes2rec.data2 {
			fmt.Printf("close connection")
			defer serverConnection.Close()
			return
		} else {
			fmt.Printf("%v\n", mes2rec)
		}
	}
}
