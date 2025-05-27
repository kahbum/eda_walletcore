#!/bin/sh

/etc/confluent/docker/run &

while ! kafka-topics --bootstrap-server localhost:9092 --list; do
  echo "Waiting for kafka to start"
  sleep 1
done

kafka-topics --bootstrap-server localhost:9092 --topic transactions --create

kafka-topics --bootstrap-server localhost:9092 --topic balances --create

wait