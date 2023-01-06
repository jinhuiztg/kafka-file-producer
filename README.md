Kafka file producer
===

Produce the content of a file to Kafka.

Build
---
```
go build -ldflags="-s -w"
```
and it'll create a executable `kafka-file-producer`.

Usage
---
```
$ ./kafka-file-producer -h
Usage of ./kafka-file-producer:
  -bootstrap-server string
        Kafka bootstrap server
  -file-path string
        File to be produced
  -topic string
        Kafka topic
```

Example
---
```
./kafka-file-producer --bootstrap-server localhost:9092 --file-path /path/to/file --topic test
```
