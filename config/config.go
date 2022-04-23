package config

import (
	"log"
	"strings"
)

var appBrokers []string
var appTopic string
var appFilePath string

func Config(topic, brokers, filePath string) {
	if brokers == "" {
		log.Fatal("--brokers not found")
	}
	if topic == "" {
		log.Fatal("--topic not found")
	}
	if filePath == "" {
		log.Fatal("--file not found")
	}

	setBrokers(brokers)
	appTopic = topic
	appFilePath = filePath

	log.Default().Printf("INFO - Kafka topic: %s\n", appTopic)
	log.Default().Printf("INFO - Kafka brokers: %s\n", appBrokers)
	log.Default().Printf("INFO - File to proccess: %s\n", appFilePath)
}

func setBrokers(brokers string) {
	appBrokers = strings.Split(brokers, ",")
}

func GetTopic() string {
	return appTopic
}

func GetBrokers() []string {
	return appBrokers
}

func GetFilePath() string {
	return appFilePath
}
