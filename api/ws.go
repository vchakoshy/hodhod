package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var wsUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketConnection struct {
	Socket *websocket.Conn
	lock   sync.Mutex
}

var clients = make(map[int64]*WebSocketConnection, 0)
var syncer = &sync.Mutex{}

func wsSendMessageClient(t int, u int64, data []byte) {
	if tc, ok := clients[u]; ok {
		tc.lock.Lock()
		defer tc.lock.Unlock()
		if err := tc.Socket.WriteMessage(t, data); err != nil {
			log.Println(err.Error())
		}
	}
}

func wsHandler(c *gin.Context, w http.ResponseWriter, r *http.Request) {
	wsUpgrade.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsUpgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	for {
		t, m, err := conn.ReadMessage()
		if err != nil {
			break
		}

		log.Printf("message type: %d, message: %s", t, m)

	}

}
