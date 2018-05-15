package db

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

func OpenRethink(prod bool) *r.Session {
	url := "172.20.10.13:28015"
	if prod {
		url = "192.168.9.95:28015"
	}

	session, err := r.Connect(r.ConnectOpts{
		Address:    url,
		InitialCap: 10,
		MaxOpen:    10,
		Database:   "multilangdb",
		Username:   "admin",
		Password:   "root-password",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	return session
}
