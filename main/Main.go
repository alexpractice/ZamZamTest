package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	configFilePath    = "../resources/config.json"
	config            = LoadConfiguration(configFilePath)
	sourceValueString = os.Args[1]
	repeatInt         int64
	channel           chan string
	wsLock            = sync.Mutex{}
)

func main() {
	verifyInputInt(sourceValueString)
	wsPort := config.WsPort
	repeatInt, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}
	channel = make(chan string, repeatInt)
	go computingAndSaving(sourceValueString)

	http.Handle("/", websocket.Handler(echo))
	if err := http.ListenAndServe(wsPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}




