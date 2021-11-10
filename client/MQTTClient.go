package client

import (
	"demo_vernemq_golang/configs"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/ztrue/shutdown"
	"sync"
	//"time"
)

var client MQTT.Client = nil;
var lock = &sync.Mutex{};

var f MQTT.MessageHandler = func (client MQTT.Client ,msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func GetMQTTClientInstance() MQTT.Client {
	if client == nil {
		lock.Lock();
		defer lock.Unlock();
		if client == nil {
			fmt.Println("Creating single instance of connection to mqtt broker");
			client = getMQTTClient()
		}
	}
	return client;
}

func getMQTTClient() MQTT.Client {
	c := MQTT.NewClient(getClientOptions())
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	addShotDownHook(c)
	return c
}

func getClientOptions() *MQTT.ClientOptions {

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", configs.MQTTBroker, configs.MQTTPort));
	opts.SetClientID("go-simple")
	//opts.SetTraceLevel(MQTT.Off);
	opts.SetDefaultPublishHandler(f);
	return opts
}

func addShotDownHook(c MQTT.Client)  {
	shutdown.Add(func() {
		fmt.Println("Killing MQTT Connection")
		c.Disconnect(250)
		fmt.Println("Killed MQTT Connection")
	})
}