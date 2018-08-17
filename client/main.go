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
	data1 int
	data2 string
}

type rmessage struct {
	data1 string
	data2 bool
}

func main() {
	mc := myClient{ip: "111"}
	serverConnection, _, err := mc.d.Dial(mc.ip, nil)
	if err != nil {
		fmt.Printf("error making connection to server")
		return
	}
	mes2send := message{0, "hello,server"}
	for {
		mes2send.data1++
		send, err := json.Marshal(mes2send)
		if err != nil {
			fmt.Printf("cannot marshal message (client to server)")
			return
		}
		err = serverConnection.WriteMessage(3, send)
		if err != nil {
			fmt.Printf("cannot send message to server")
			return
		}
		_, receive, err := serverConnection.ReadMessage()
		if err != nil {
			fmt.Printf("cannot receive message from server")
			return
		}
		var mes2rec rmessage
		err = json.Unmarshal(receive, &mes2rec)
		if err != nil {
			fmt.Printf("cannot Unmarshal message from server")
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
