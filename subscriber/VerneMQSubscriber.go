package subscriber

import (
	"demo_vernemq_golang/client"
	"fmt"
	"os"
)

func Subscribe(topic string) {
	c := client.GetMQTTClientInstance();
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func Unsubscribe(topic string) {
	c := client.GetMQTTClientInstance();
	if token := c.Unsubscribe(topic	); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
