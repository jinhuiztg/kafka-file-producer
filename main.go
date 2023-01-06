package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {
	bootstrapServer := flag.String("bootstrap-server", "", "Kafka bootstrap server")
	topic := flag.String("topic", "", "Kafka topic")
	filePath := flag.String("file-path", "", "File to be produced")
	flag.Parse()

	if *bootstrapServer == "" || *topic == "" || *filePath == "" {
		fmt.Println("--bootstrap-server, --topic, and --filepath must be provided.")
		fmt.Println("Usage of producer:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	producer, err := sarama.NewSyncProducer(strings.Split(*bootstrapServer, ","), nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Failed to read file: %v", err)
		os.Exit(1)
	}

	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

}
