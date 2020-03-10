package mqtt

import (
	"fmt"
	"log"
	"net/url"
)

const (
  protocol = "tcp"
)

func CreateURL(protocol string, ipAddress string, port int16, topic string) url.URL {
  uri, err := url.Parse(fmt.Sprintf("%v://%v:%v/%v", protocol, ipAddress, port, topic))
	if err != nil {
		log.Fatal(err)
	}
  return *uri
}

