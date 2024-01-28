# kafka-with-go

# Get the Kafka

For example 

$ tar -xzf kafka_2.13-3.6.1.tgz
$ cd kafka_2.13-3.6.1

# start zookeeper in a terminal
./bin/zookeeper-server-start.sh config/zookeeper.properties

# Start Kafka Server in a terminal
./bin/kafka-server-start.sh config/server.properties

# Create a topic in a terminal
./bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server 192.168.49.1:9092

# Produce the message in a terminal
./bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server 192.168.49.1:9092

# Consumer the produced message in a terminal
./bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server 192.168.49.1:9092
