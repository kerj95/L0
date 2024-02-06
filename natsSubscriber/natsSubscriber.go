package natsSubscriber

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func NatsSubscriber() {
	clusterID := "prod"
	clientID := "l0"
	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()
	for {
		sc.Subscribe("foo", func(m *stan.Msg) {
			fmt.Printf("Received a message: %s\n", string(m.Data))
		}, stan.StartWithLastReceived(), stan.DurableName("test"))
	}
}
