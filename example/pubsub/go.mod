module pubsub

go 1.23.4

replace github.com/cans-communication/hermes => ../..

require (
	github.com/cans-communication/hermes v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/kelseyhightower/envconfig v1.4.0
)

require (
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/nats-io/nats.go v1.39.1 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
