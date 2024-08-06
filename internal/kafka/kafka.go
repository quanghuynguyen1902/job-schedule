package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Kafka struct {
	brokers  string
	producer *producer
}

type producer struct {
	producer *kafka.Producer
	ready    bool
	termChan chan bool
}

func New(brokers string) *Kafka {
	kafka := &Kafka{
		brokers:  brokers,
		producer: &producer{},
	}

	return kafka
}
