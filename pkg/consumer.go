package pkg

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"

	"log"

	"consumer/test/config"
)

type TopicStruct struct {
	Header interface{} `json:"header"`
	Body   string     `json:"body"`
}

func Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.KafkaBroker},
		GroupID:  config.Topic+"-01",
		Topic:    config.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for{
		var consumedData TopicStruct
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err.Error())
			continue
		}
		err = json.Unmarshal([]byte(string(m.Value)), &consumedData)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println(consumedData.Body)

	}
	r.Close()
}