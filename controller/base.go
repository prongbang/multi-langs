package controller

import (
	"github.com/labstack/echo"
	r "gopkg.in/gorethink/gorethink.v4"
)

type Controller struct {
	Ctx  echo.Context
	RTDb *r.Session
}
