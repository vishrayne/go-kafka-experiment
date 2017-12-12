# go-kafka-experiment
Attempt to get familiar with [kafka](https://kafka.apache.org/intro) using golang. 
For learning purpose, I'm going with confluent's platform and corresponding golang client.

# setup
1. Install kafka (Apache version or Confluent's)
2. Install [librdkafka](https://github.com/edenhill/librdkafka)
3. Start a broker/zookeeper [Assuming this to be running on localhost:9092 (default=9092)]
```
$ cd go-kafka-sample
$ dep ensure
$ go [install|build] github.com/vishrayne/go-kafka-sample
$ go-kafka-sample 
```