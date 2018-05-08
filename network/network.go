package network

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the router.
	pingWait = 90 * time.Second

	// Interval to wait between retry attempts
	retryInterval = 3 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwitchConfigSet struct {
	Name   string   `json:"name"`
	Config []Config `json:"config"`
}

type Config struct {
	Multicast `json:"multicast"`
	Name      string `json:"name"`
}

type Multicast struct {
	Address string `json:"address"`
}

type Command struct {
	Creds
	SwitchConfigSet `json:"config_set"`
}

func OpenConnection() error {
	// Building JSON Query
	Fig := []Config{Config{Multicast: Multicast{Address: "239.0.0.1"}, Name: "ip_input1"}, Config{Multicast: Multicast{Address: "239.10.0.1"}, Name: "ip_input3"}}
	SC := Command{Creds: Creds{Username: "admin", Password: "password"}, SwitchConfigSet: SwitchConfigSet{Name: "ip_input", Config: Fig}}
	fmt.Println(SC)
	Comm, err := json.Marshal(SC)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output to Console the factored JSON
	m := string(Comm)
	fmt.Println(m)

	// Open connection to the decoder
	dialer := &websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.Dial(fmt.Sprintf("ws://192.168.0.7/wsapp/"), nil)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem establishing the websocket with 192.168.0.7: %v", err.Error()))
		return err
	}

	// Output to console the connection information
	//if conn != nil {
	//	fmt.Println(conn)
	//	fmt.Println(test)
	//}

	// Write JSON to Connected Decoder
	err = conn.WriteMessage(websocket.TextMessage, Comm)
	if err != nil {
		log.Println(err)
		return (err)
	}

	// Read Back any message that is returned from Writing the Message
	_, msgd, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	msg := string(msgd)
	fmt.Println(msg)
	conn.Close()
	return nil
}
