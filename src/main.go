package main

import (
	"fmt"
	"sync"

	// "goflush/src/core"
	"goflush/src/utils"
	"goflush/src/wsnet"
)

func main() {
	var wg sync.WaitGroup

	// This is an incoming websocket message queue
	var messageQueue utils.Queue[[]byte]

	// Start queue consumer in a goroutine
	wg.Add(1)
	go messageQueue.Watch(&wg, func(msg []byte) {
		fmt.Println("From callback", msg)
	})

	// Initiate server
	wsnet.InitiateWebsocketServer(&messageQueue)
}
