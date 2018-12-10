package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"os"
	"time"
)

//функция для выдачи сообщений клиенту, подключившемуся по WebSocket
func echo(ws *websocket.Conn) {
	count := repeatInt + 1
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
			close(channel)
			break
		}
		time.Sleep(3 * time.Second)
	}
	os.Exit(3)
}
