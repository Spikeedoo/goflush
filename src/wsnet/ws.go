package wsnet

import (
	// "crypto/sha1"
	// "encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"goflush/src/utils"
)

// func handleWebsocketMessage(r *http.Request, conn *websocket.Conn, messageType int, message []byte) {
// 	// Hash the incoming IP as the client's identifier
// 	incomingIpAddress := r.RemoteAddr
// 	h := sha1.New()
// 	h.Write([]byte(incomingIpAddress))
// 	clientId := hex.EncodeToString(h.Sum(nil))
// 	fmt.Println(clientId)
// }

func wsEndpoint(messageQueue *utils.Queue[[]byte]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // TODO: Revisit this for production
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading to WebSocket connection!")
			return
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}

			messageQueue.Push(message)
		}
	}
}

func InitiateWebsocketServer(messageQueue *utils.Queue[[]byte]) {
	http.HandleFunc("/ws", wsEndpoint(messageQueue))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
