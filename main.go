package main

import (
	"L0/natsSubscriber"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	go natsSubscriber.NatsSubscriber()
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
