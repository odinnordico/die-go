package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	messaging "./communication/messaging/mqtt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func listen(uri *url.URL, topic string) {
	client := messaging.Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	uri, err := url.Parse("mqtt://127.0.0.1:1883/test")
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	fmt.Printf("%s: %v\n", "uri", uri)
	if topic == "" {
		topic = "test"
	}

	go listen(uri, topic)

	client := messaging.Connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		messaging.Publish(client, topic, 0, false, "{msg="+t.String()+"}")
	}
}
