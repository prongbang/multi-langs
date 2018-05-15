package handler

import (
	"github.com/labstack/echo"
	r "gopkg.in/gorethink/gorethink.v4"
)

type Handler struct {
	RTDb *r.Session
}

func (handle Handler) NewChangesHandler(fn func(chan interface{}), c echo.Context) error {
	h := newHub()
	go h.run()

	fn(h.broadcast)

	return WsHandler(h, c)
}
