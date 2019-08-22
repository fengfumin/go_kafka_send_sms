package kafka_sms_consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)


// 支持brokers cluster的消费者
func (s *Server) StartHandleSmsConsumer(brokers []string, topics string) (err error) {
	return s.handlerSmsConsumer(brokers, topics)
}

func (s *Server) handlerSmsConsumer(brokers []string, topic string) (err error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return
	}
	defer consumer.Close()
	pConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest, )
	defer pConsumer.Close()
	fmt.Println("SMS Server Init Success",
		"\nBroker", brokers,
		"\nListen at Topic", topic,
		"\nStart Consuming", time.Now().Format("2006-01-02 15:04"))
	for {
		select {
		case msg := <-pConsumer.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
			err:=s.SendSms(msg.Value)
			if err!=nil{
				fmt.Println(err)
			}
		case err := <-pConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}


