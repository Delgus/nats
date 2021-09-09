package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	user := os.Getenv("NATS_USERNAME")
	psw := os.Getenv("NATS_PASSWORD")
	subject := os.Getenv("NATS_SUBJECT")

	log.Println(natsURL, user, psw, subject)

	// Connect to NATS
	nc, err := nats.Connect(natsURL, nats.UserInfo(user, psw))
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		err := nc.Publish(subject, []byte("Hello World"))
		if err != nil {
			log.Println(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	<-c
}
