package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	user := os.Getenv("NATS_USERNAME")
	psw := os.Getenv("NATS_PASSWORD")
	subject := os.Getenv("NATS_SUBJECT")
	clusterID := os.Getenv("NATS_STREAMING_CLUSTER_ID")

	conn, err := stan.Connect(
		clusterID,
		"streaming_pub",
		stan.NatsOptions(nats.UserInfo(user, psw)),
		stan.NatsURL(natsURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	var number int

	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		number++

		err = conn.Publish(subject, []byte("Hello from streaming number:"+strconv.Itoa(number)))
		if err != nil {
			log.Fatal(err)
		}
	}

	// don't forget close connection !!!
	conn.Close()
}
