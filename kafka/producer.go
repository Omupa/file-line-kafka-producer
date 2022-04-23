package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	"omupa.com/file-line-kafka-producer/config"
)

var localProducer sarama.SyncProducer

func SendMessageSync(message string) error {
	if localProducer == nil {
		newProducer()
	}

	msg := prepareMessage(message)

	partition, offset, err := localProducer.SendMessage(msg)
	if err != nil {
		log.Default().Printf("%s error occured.", err.Error())
		return err
	} else {
		log.Default().Printf("Message sended - partion: %d offset: %d\n", partition, offset)
	}

	return nil
}

func newProducer() {
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(config.GetBrokers(), producerConfig)

	if err != nil {
		log.Fatal("Kafka producer cannot be initialized", err)
	} else {
		log.Default().Println("Kafka producer initialized")
	}

	localProducer = producer
}

func prepareMessage(message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     config.GetTopic(),
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}
