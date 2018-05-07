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
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type command struct {
	username []string
	password []string
	addrInput []string
	addrOutput []string
}
/* Left over from previous code
type Node struct {
        Name          string
        Conn          *websocket.Conn
        WriteQueue    chan base.Message
        ReadQueue     chan base.Message
        DecoderAddress string
        filters       map[string]bool
        readDone      chan bool
        writeDone     chan bool
        lastPingTime  time.Time
        state         string
}
*/
//func (n *Node) OpenConnection() error {
func OpenConnection() error {
        //var mess = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.1"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.1"},"name": "ip_input3"}]}}`)
	
        var m = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.2"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.2"},"name": "ip_input3"}]}}`)

        //var mess = []byte(`{"username":"admin","password":"Atlona","config_set":{"name":"ip_input","config":[{"multicast": {"address": "239.1.1.11"},"name": "ip_input1"},{"multicast": {"address": "239.10.1.11"},"name": "ip_input3"}]}}`)
	//open connection to the decoder
        dialer := &websocket.Dialer{
                HandshakeTimeout: 10 * time.Second,
        }

        //conn, _, err := dialer.Dial(fmt.Sprintf("ws://%s/wsapp", n.DecoderAddress), nil)
        conn, test, err := dialer.Dial(fmt.Sprintf("ws://192.168.0.7/wsapp/"), nil)
	if err != nil {
                log.Printf(color.HiRedString("There was a problem establishing the websocket with 192.168.0.7: %v", err.Error()))
                return err
        }
        //n.Conn = conn
        if conn != nil {
		fmt.Println(conn)
		fmt.Println(test) 
	}
	err = conn.WriteMessage(websocket.TextMessage, m)
	if err != nil {
		log.Println(err)
		return(err)
	}
	conn.Close()
        return nil
}
