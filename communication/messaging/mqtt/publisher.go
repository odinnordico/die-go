package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(client mqtt.Client, topic string, qualityOfService byte, retained bool, payload interface{}) {
	client.Publish(topic, qualityOfService, retained, payload)
}
