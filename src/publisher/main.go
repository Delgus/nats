package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
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

	defer nc.Close()

	var number int

	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		number++

		err = nc.Publish(subject, []byte("nats_message_"+strconv.Itoa(number)))
		if err != nil {
			log.Fatal(err)
		}
	}
}
