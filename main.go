package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"syscall"
)

var (
	projectID = pflag.StringP("projectID", "p", "", "project ID")
	topic     = pflag.StringP("topic", "t", "", "topic")
	subName   = pflag.StringP("sub", "s", "test-sub", "topic")
)

func main() {
	pflag.Parse()
	if *projectID == "" {
		exit(errors.New("project ID is not provided. use --projectID to set"), 1)
	}
	if *topic == "" {
		exit(errors.New("topic is not provided. use --topic to set"), 1)
	}
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		exit(errors.New("PUBSUB_EMULATOR_HOST environment variable is not set"), 1)
	}

	psClient, err := pubsub.NewClient(context.Background(), "test")
	if err != nil {
		exit(err, 1)
	}
	defer psClient.Close()

	subConfig := pubsub.SubscriptionConfig{
		Topic: psClient.Topic(*topic),
	}

	sub, err := psClient.CreateSubscription(context.Background(), *subName, subConfig)
	if err != nil {
		exit(err, 1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	sub.ReceiveSettings.NumGoroutines = 1
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			fmt.Println(string(m.Data))
			m.Ack()
		})
		if err != nil {
			exit(err, 1)
		}
	}()

	<-sigChan
	cancel()
	exit(errors.New("terminated"), 0)
}

func exit(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
