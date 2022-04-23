package file

import (
	"bufio"
	"log"
	"os"

	"omupa.com/file-line-kafka-producer/config"
	"omupa.com/file-line-kafka-producer/kafka"
)

func ReadFileAndSendMessage() {
	messageAmoung := 1

	file, err := os.Open(config.GetFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		messageAmoung++

		err := kafka.SendMessageSync(scanner.Text())

		if err != nil {
			messageAmoung--
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Default().Printf("Proccess finished. Messages sended: %d", messageAmoung)
	}
}
