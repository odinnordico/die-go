package main

import (
	"time"

	messaging "./connection/messaging/mqtt"
        conn "./connection"
)

const (
  topic string = "test"
  ipAddress string = "127.0.0.1"
  port int16 = 1883
  protocol string = "mqtt"
)


func main() {
  uri := conn.CreateURL(protocol, ipAddress, port, topic)
	go messaging.Listen("subscriber", &uri, topic, nil)

	client := messaging.Connect("pub", &uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		messaging.Publish(client, topic, 0, false, "{msg="+t.String()+"}")
	}
}
