# Kafka Sample Project

This is a sample Kafka project using Docker Compose and Go.

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Go

### Setup and Run

#### Start Kafka Services
To start the Kafka and Zookeeper containers, run:
```sh
$ make up
```

#### Stop kafka Services

To stop and remove the containers, run:
```sh
make down
```

#### Run the Publisher
To send messages to Kafka, run
```sh
make publisher
```

#### Run the Consumer
To consume messages from Kafka, run
```sh
make consumer
```

### Project Structure

```sh
.
├── cmd
│   ├── publisher   # Kafka producer (message sender)
│   ├── consumer    # Kafka consumer (message receiver)
├── docker-compose.yml # Kafka & Zookeeper configuration
├── Makefile         # Shortcut commands for convenience
└── README.md        # Project documentation
```


Notes
- The default Kafka broker address is localhost:9093
- The topic used in this sample is "post-likes"
- Modify the consumer and publisher logic as needed

