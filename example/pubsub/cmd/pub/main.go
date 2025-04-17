package main

import (
	"fmt"

	"github.com/cans-communication/hermes"
	"github.com/kelseyhightower/envconfig"
)

type EnvCfg struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Subject  string `envconfig:"SUBJECT" required:"true"`
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
	err = p.ProduceAsync(
		cfg.Subject,
		[]byte(`{"id":"1", "name": "alice"}`),
	)

	if err != nil {
		panic(err)
	}

}
