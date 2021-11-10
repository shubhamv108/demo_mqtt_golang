package publisher

import (
	"demo_vernemq_golang/client"
	_ "github.com/eclipse/paho.mqtt.golang"
)

func Publish(topic string, text string) {
    c := client.GetMQTTClientInstance()
	token := c.Publish(topic, 0, false, text)
	token.Wait()
	//for i := 0; i < 5; i++ {
	//	text := fmt.Sprintf("this is msg #%d!", i)
	//	token := c.Publish("go-mqtt/sample", 0, false, text)
	//	token.Wait()
	//}

}