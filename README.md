# GoExpert-Events

GoExpert-Events is a Go-based project that manages events — queues, producers and consumers. The repository includes the necessary files and folders to set up and run the application.

## Getting Started

Build and run the application using Docker Compose:

```sh
docker-compose up -d
```

That will bootstrap RabbitMQ and its management dashboard. 

Access http://localhost:15672 with username `guest` and pass `guest` to create a queue named `my-queue`.

Finally, run the producer and consumer to see the events being managed by the application.

```sh
# Start producer
cd cmd/producer
go run main.go

# Start consumer
cd cmd/consumer
go run main.go
```
