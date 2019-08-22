package kafka_sms_producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type KafkaMsgProducer struct {
	Addr   []string
	Config *sarama.Config
}

func NewKafkaMsgProducer(Addr []string, Config ...*sarama.Config) *KafkaMsgProducer {
	var config *sarama.Config
	if len(Config) == 0 {
		config = sarama.NewConfig()
	} else {
		config = Config[0]
	}
	return &KafkaMsgProducer{
		Addr,
		config,
	}
}

func (p *KafkaMsgProducer) SendMsgToTopic(topic, value string) {
	producer, e := sarama.NewAsyncProducer(p.Addr, p.Config)
	if e != nil {
		fmt.Println(e)
		return
	}

	defer producer.AsyncClose()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	go func(p sarama.AsyncProducer) {
			select {
			case suc := <-producer.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-producer.Errors():
				fmt.Println("err: ", fail.Err)
			default:
				log.Println("Produced message default", )
			}
	}(producer)

	producer.Input()<- msg
	time.Sleep(1 * time.Nanosecond)
}
