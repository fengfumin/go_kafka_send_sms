package main

import (
	"awesomeProject/kafka_sms_producer"
)

func main() {

	p:=kafka_sms_producer.NewKafkaMsgProducer([]string{"10.0.0.11:9092"})

	p.SendMsgToTopic("sms",`{"phone":"电话号码","content":"90909"}`)

}
