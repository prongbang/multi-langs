package handler

import (
	"log"
	"multi-langs/utils"

	r "gopkg.in/gorethink/gorethink.v4"
)

func (handle Handler) AttributesChanges(ch chan interface{}) {
	// Use goroutine to wait for changes. Prints the first 10 results
	go func() {
		for {
			res, err := r.Table(utils.TABLE_ATTRIBUTES).Changes().Run(handle.RTDb)
			if err != nil {
				log.Fatalln(err)
			}

			var response interface{}
			for res.Next(&response) {
				ch <- response
			}

			if res.Err() != nil {
				log.Println(res.Err())
			}
		}
	}()
}

// func activeChanges(ch chan interface{}) {
// 	// Use goroutine to wait for changes. Prints the first 10 results
// 	go func() {
// 		for {
// 			res, err := r.DB("todo").Table("items").Filter(r.Row.Field("Status").Eq("active")).Changes().Run(session)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}

// 			var response interface{}
// 			for res.Next(&response) {
// 				ch <- response
// 			}

// 			if res.Err() != nil {
// 				log.Println(res.Err())
// 			}
// 		}
// 	}()
// }
// func completedChanges(ch chan interface{}) {
// 	// Use goroutine to wait for changes. Prints the first 10 results
// 	go func() {
// 		for {
// 			res, err := r.DB("todo").Table("items").Filter(r.Row.Field("Status").Eq("complete")).Changes().Run(session)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}

// 			var response interface{}
// 			for res.Next(&response) {
// 				ch <- response
// 			}

// 			if res.Err() != nil {
// 				log.Println(res.Err())
// 			}
// 		}
// 	}()
// }
