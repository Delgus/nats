package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	user := os.Getenv("NATS_USERNAME")
	psw := os.Getenv("NATS_PASSWORD")
	subject := os.Getenv("NATS_SUBJECT")

	nc, err := nats.Connect(natsURL, nats.UserInfo(user, psw))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("sending message")

	err = nc.Publish(subject, []byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}

	// don't forget close connection !!!
	nc.Close()
}
