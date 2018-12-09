package main

import (
	"github.com/mediocregopher/radix.v2/pool"
	"log"
)

var (
	db     *pool.Pool
	err    error
	dbHost = config.DbConfig.Host
	dbPort = config.DbConfig.Port
)

func init() {
	db, err = pool.New("tcp", dbHost+dbPort, 10)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

func saveToDataBase(key string, value string) {
	dbClientConn, err := db.Get()
	defer db.Put(dbClientConn)
	if err != nil {
		log.Panic(err)
	}
	dbClientConn.Cmd("RPUSH", key, value)
}
