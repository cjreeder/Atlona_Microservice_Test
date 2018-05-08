package network

import (
	"fmt"
	"log"
	"time"

	"github.com/byuoitav/event-router-microservice/base"
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
	//var mess = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.1"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.1"},"name": "ip_input3"}]}}`)
	//var m = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.2"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.2"},"name": "ip_input3"}]}}`)
	//var mess = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.11"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.11"},"name": "ip_input3"}]}}`)
	less := []Config{Config{Multicast: Multicast{Address: "239.0.0.1"}, Name: "ip_input1"}, Config{Multicast: Multicast{Address: "239.10.0.1"}, Name: "ip_input3"}}
	Test := Command{Creds: Creds{Username: "admin", Password: "password"}, SwitchConfigSet: SwitchConfigSet{Name: "ip_input", Config: less}}
	fmt.Println(Test)
	Foo, err := json.Marshal(Test)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//Bar, err := json.Marshal(Command{Creds: Creds{Username: "admin", Password: "password"}, ConfigSet: ConfigSet{Name: "ip_table", Config: Config{Multicast: Multicast{Address:"239.0.0.1"}}}})
	if err != nil {
		fmt.Println("Error:", err)
	}
	m := string(Foo)

	//open connection to the decoder
	dialer := &websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, test, err := dialer.Dial(fmt.Sprintf("ws://192.168.0.7/wsapp/"), nil)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem establishing the websocket with 192.168.0.7: %v", err.Error()))
		return err
	}
	if conn != nil {
		fmt.Println(conn)
		fmt.Println(test)
	}
	err = conn.WriteMessage(websocket.TextMessage, m)
	if err != nil {
		log.Println(err)
		return (err)
	}
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m)
	conn.Close()
	return nil
}
