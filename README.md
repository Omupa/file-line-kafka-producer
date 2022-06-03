# file-line-kafka-producer
CLI to read file line by line and send this line to Kafka topic

# Build
Firstly you need to install Golang an then you will be able to run this command `go build`.

After a binary package will be on the path of the project.

# Run and use
With the binary it's possivel to run the program with the follow sample command
```
./file-line-kafka-producer --brokers "kafka-broker-01:9092,kafka-broker-02:9092" --topic "target-topic-name" --file "/path/to/file/line-by-line.txt"
```
