package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"os"
	"time"
)

func echo(ws *websocket.Conn) {
	count := repeatInt + 1;
	resCount := int64(0)
	wsLock.Lock()
	defer wsLock.Unlock()
	for {
		reply := <-channel
		if err = websocket.Message.Send(ws, reply); err != nil {
			fmt.Println("Can't send")
			break
		}
		resCount++
		if count == resCount {
			ws.Close()
			break
		}
		time.Sleep(3 * time.Second)
	}
	os.Exit(3)
}