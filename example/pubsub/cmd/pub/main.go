package main

import (
	"fmt"
	"time"

	"github.com/cans-communication/hermes"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type EnvCfg struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Subject  string `envconfig:"SUBJECT" required:"true"`
}

type Message struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
	Msg       string    `json:"msg"`
}

func main() {
	fmt.Println("publisher start!")

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

	p := client.Producer()
	err = p.ProduceAsyncJson(
		cfg.Subject,
		&Message{
			Timestamp: time.Now(),
			ID:        uuid.New().String(),
			Msg:       "Hello",
		},
	)

	if err != nil {
		panic(err)
	}

}
