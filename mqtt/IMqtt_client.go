package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQClient CLIENT
var MQClient IMQTTClient

// IMQTTClient - IModelRepository
type IMQTTClient interface {
	//Set - Set a value with key to Redis DB
	CreateMQTTClient() error
	Start() (bool, error)
	Stop() error
	Subscribe(topic string, handler mqtt.MessageHandler, qos byte) error
	//SubscribeOC(topic string, handler mqtt.MessageHandler, qos byte) *mqtt.ClientOptions
	SubscribeD(topic string) error
	UnSubscribe(topic string) error
	Publish(topic, message string) error
	PublishQOS(topic, message string, qos byte) error
}