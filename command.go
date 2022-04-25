package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"omupa.com/file-line-kafka-producer/config"
	"omupa.com/file-line-kafka-producer/file"
)

// ./file-line-kafka-producer run --brokers "kafka-broker.url:9092" --topic "topic-name" --file "/path/to/file.txt"

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "file-line-kafka-producer"
	app.Description = "This CLI read file by line and send this line to Kafka topic"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "brokers",
			Usage: "Kafka brokers separated by comma.",
		},
		&cli.StringFlag{
			Name:  "topic",
			Usage: "Topic to send data read from file line",
		},
		&cli.StringFlag{
			Name:  "file",
			Usage: "File path on disk",
		},
	}
	app.Action = func(c *cli.Context) error {
		brokers := c.String("brokers")
		topic := c.String("topic")
		filePath := c.String("file")

		config.Config(topic, brokers, filePath)

		file.ReadFileAndSendMessage()

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
