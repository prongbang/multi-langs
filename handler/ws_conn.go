package handler

import (
	"log"

	"github.com/labstack/echo"

	"github.com/gorilla/websocket"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan interface{}
}

func (c *connection) reader() {
	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for change := range c.send {
		err := c.ws.WriteJSON(change)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

// func WsHandler(h hub) http.HandlerFunc {
// 	log.Println("Starting websocket server")
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ws, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			return
// 		}
// 		c := &connection{send: make(chan interface{}, 256), ws: ws}
// 		h.register <- c
// 		defer func() { h.unregister <- c }()
// 		go c.writer()
// 		c.reader()
// 	}
// }

func WsHandler(h hub, c echo.Context) error {
	log.Println("Starting websocket server")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	// defer ws.Close()
	conn := &connection{send: make(chan interface{}, 256), ws: ws}
	h.register <- conn
	defer func() { h.unregister <- conn }()
	go conn.writer()
	conn.reader()

	return nil
}
