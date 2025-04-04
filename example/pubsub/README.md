## Publish Subscriber Example


## Env
```sh
HERMES_EXAMPLE_HOST=localhost
HERMES_EXAMPLE_PORT=4222
HERMES_EXAMPLE_USER=dev
HERMES_EXAMPLE_PASSWORD=devpassword
HERMES_EXAMPLE_SUBJECT=local-nats.message
HERMES_EXAMPLE_STREAM=local-nats
HERMES_EXAMPLE_CONSUMER=local-nats-consumer
```

## Setup NATs
1. compose up
```sh
docker compose up -d
```
2. set up nats
2.1. add stream
```sh
nats stream add --server $HERMES_EXAMPLE_HOST:$HERMES_EXAMPLE_PORT --user $HERMES_EXAMPLE_USER --password $HERMES_EXAMPLE_PASSWORD --storage=memory --subjects=$HERMES_EXAMPLE_SUBJECT --replicas=1  --retention=limits --discard=old --max-msgs=-1 --max-msgs-per-subject=-1 --no-allow-rollup --max-bytes=-1 --max-age=1h --max-msg-size=-1 --dupe-window=2m --deny-delete --deny-purge $HERMES_EXAMPLE_STREAM 
```

2.2 add consumer
```sh
nats consumer add --server $HERMES_EXAMPLE_HOST:$HERMES_EXAMPLE_PORT --user $HERMES_EXAMPLE_USER --password $HERMES_EXAMPLE_PASSWORD --deliver=all --pull --ack=all --replay=instant --max-deliver=-1 --max-pending=0 --no-headers-only --backoff=none --filter=$HERMES_EXAMPLE_SUBJECT $HERMES_EXAMPLE_STREAM $HERMES_EXAMPLE_CONSUMER 
```

2.3 info stream
```sh
nats stream info --server $HERMES_EXAMPLE_HOST:$HERMES_EXAMPLE_PORT --user $HERMES_EXAMPLE_USER --password $HERMES_EXAMPLE_PASSWORD
```

2.4 info consumer
```sh
nats consumer info --server $HERMES_EXAMPLE_HOST:$HERMES_EXAMPLE_PORT --user $HERMES_EXAMPLE_USER --password $HERMES_EXAMPLE_PASSWORD
```

## How to run ?

1. export environment variable
```sh
export $(grep -v '^#' .env | xargs)
```
2. run publish
```sh
go run ./cmd/pub
```
