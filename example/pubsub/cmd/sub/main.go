package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/cans-communication/hermes"
	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go/jetstream"
)

type EnvCfg struct {
	Host       string `envconfig:"HOST" required:"true"`
	Port       int    `envconfig:"PORT" required:"true"`
	User       string `envconfig:"USER" required:"true"`
	Password   string `envconfig:"PASSWORD" required:"true"`
	Stream     string `envconfig:"STREAM" required:"true"`
	ConsumerID string `envconfig:"CONSUMER" required:"true"`
	Subject    string `envconfig:"SUBJECT" required:"true"`
}

type Message struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
	Msg       string    `json:"msg"`
}

func main() {

	var cfg EnvCfg
	err := envconfig.Process("HERMES_EXAMPLE", &cfg)
	if err != nil {
		panic(err)
	}

	client, err := hermes.Connect(
		hermes.ConnectOpt{
			Host:     cfg.Host,
			Port:     cfg.Port,
			User:     cfg.User,
			Password: cfg.Password,
		},
	)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	ctx := context.Background()
	c, err := client.Consumer(
		ctx,
		cfg.Stream,
		cfg.ConsumerID,
	)
	if err != nil {
		panic(err)
	}

	sub, err := c.Consume(func(msg jetstream.Msg) {
		msg.Ack()

    var payload Message 
    err := json.Unmarshal(msg.Data(), &payload)
    if err != nil {
      return
    }

    d, err := json.MarshalIndent(payload, " ", " ")
    if err != nil {
      return
    }

		fmt.Println(string(d))

	})

	if err != nil {
		panic(err)
	}

	nctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	<-nctx.Done()

	sub.Stop()
}
