package main

import (
	"GoKafkaMessenger/internal/infra/akafka"
	"GoKafkaMessenger/internal/infra/repository"
	"GoKafkaMessenger/internal/usecase"
	"database/sql"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *kafka.Message)

	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repo := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repo)

	for msg := range msgChan {

		dto := usecase.CreateProductInputDto{}

		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Println("error on unmarshal value: ", err.Error())
			continue
		}

		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			log.Println("error on execute use case function: ", err.Error())
			continue
		}
	}

}
