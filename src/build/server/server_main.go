package main

import (
	"fmt"
	"log"

	"db"
	"server"
)

var svr *server.Server

func main() {
	fmt.Println("dreamdb run--------------------")
	cache, err := db.Open("dream.db", 0644)
	if err != nil {
		log.Fatalf("db.Open error %s", err.Error())
	}

	err = createServer(cache)
	if err != nil {
		log.Fatalf("createServer error %s", err.Error())
	}

	err = runServer()
	if err != nil {
		log.Fatalf("run server error %s", err.Error())
	}
}

func createServer(cache *db.DB) (err error) {
	svr, err = server.NewServer(cache)
	return
}

func runServer() (err error) {
	err = svr.Run()
	return
}
