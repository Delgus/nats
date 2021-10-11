package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
		"streaming_sub",
		stan.NatsOptions(nats.UserInfo(user, psw)),
		stan.NatsURL(natsURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	sub, err := conn.Subscribe(subject, func(msg *stan.Msg) {
		log.Printf("get data  %q \n", msg.Data)
	})
	if err != nil {
		log.Fatal(err)
	}

	defer sub.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	<-c
}