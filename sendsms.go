package main

import (
	"awesomeProject/kafka_sms_producer"
)

func main() {

	p:=kafka_sms_producer.NewKafkaMsgProducer([]string{"10.0.0.11:9092"})

	p.SendMsgToTopic("sms",`{"phone":"13916895160","content":"90909"}`)

}
