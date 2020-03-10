 package mqtt

import (
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func handleMessage (client mqtt.Client, msg mqtt.Message) {
  fmt.Printf("[%v] => %v\n", msg.Topic(), string(msg.Payload()))
}

func Listen(clientID string, uri *url.URL, topic string, callback mqtt.MessageHandler) {
	client := Connect(clientID, uri)
        messageHandler:= handleMessage
        if callback != nil {
          messageHandler = callback
        }
	client.Subscribe(topic, 0, messageHandler)
}
