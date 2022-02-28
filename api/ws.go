package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
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

var clients = make(map[uint64]*WebSocketConnection, 0)
var syncer = &sync.Mutex{}

func wsSendMessageClient(t int, u uint64, data []byte) {
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

	log.Println("new connection established")

	for {
		t, m, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if strings.HasPrefix(string(m), "login: ") {
			username := "demo"
			pass := "demo"
			// TODO: use interface later
			ad, err := AuthenticationDummy{}.Login(username, pass)
			if err != nil {
				conn.WriteMessage(1, []byte(err.Error()))
				continue
			}
			ws := new(WebSocketConnection)
			ws.Socket = conn
			clients[ad.UserID] = ws
			conn.WriteMessage(1, []byte(fmt.Sprintf("Login successful %s", username)))
			continue
		}

		log.Printf("message type: %d, message: %s", t, m)

	}

}
