package main

import (
	"GoKafkaMessenger/internal/infra/akafka"
	"GoKafkaMessenger/internal/infra/repository"
	"database/sql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *kafka.Message)

	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewProductRepositoryMysql(db)

}
