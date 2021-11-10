package main

import (
	VerneMQPublisher "demo_vernemq_golang/publisher"
	VerneMQSubsciber "demo_vernemq_golang/subscriber"
	"fmt"
	"github.com/ztrue/shutdown"
	"time"
)


func main() {
	topic :=  "go-mqtt/sample"

	VerneMQSubsciber.Subscribe(topic)

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		VerneMQPublisher.Publish(topic, text)
	}

	time.Sleep(3 * time.Second)

	VerneMQSubsciber.Unsubscribe(topic)

	shutdown.Listen()
}