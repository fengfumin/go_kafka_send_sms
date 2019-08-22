package main

import (
	"awesomeProject/kafka_sms_consumer"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.AddConfigPath("./kafka_sms_consumer")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	server := kafka_sms_consumer.NewSmsServer(config.GetString("keyId"),
		config.GetString("keySecret"),
		config.GetString("signName"),
		config.GetString("templateCode"))
	if err := server.StartHandleSmsConsumer(config.GetStringSlice("kafka-brokers"), config.GetString("topic")); err != nil {
		panic(err)
	}

}
