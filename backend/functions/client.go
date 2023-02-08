package functions

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var mutex = &sync.RWMutex{}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan ChatFields
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		H.unregister <- s
		c.ws.Close()
	}()
	for {
		var chatFields ChatFields
		err := c.ws.ReadJSON(&chatFields)
		chatFields.Id = s.room
		chatFields.MessageId = Generate()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			return
		}
		m := data{chatFields}
		H.broadcast <- m
		SqlExec.messages <- chatFields
	}

}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	defer func() {
		c.ws.Close()
	}()
	for {
		message, ok := <-c.send
		if !ok {
			// filter.delete(s.sessionId)
			c.ws.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := c.ws.WriteJSON(message)
		if err != nil {
			// filter.delete(s.sessionId)
			fmt.Println("error writing to chat:", err)
			return
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	id := <-chatroomId
	user := <-loggedInUsername
	fmt.Println("ws id", id)
	fmt.Println("ws opened", user)
	cookie, _ := r.Cookie("session")
	c := &connection{send: make(chan ChatFields, 1), ws: ws}
	s := subscription{c, id, user, cookie.Value}
	H.register <- s
	go s.writePump()
	go s.readPump()
}
