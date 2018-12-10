package main

import (
	"github.com/mediocregopher/radix.v2/pool"
	"log"
)

var (
	db     *pool.Pool
	err    error
	dbHost = config.DbConfig.Host //хост БД
	dbPort = config.DbConfig.Port //порт БД
)

func init() {
	db, err = pool.New("tcp", dbHost+dbPort, 10)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

//функция сохранения значения value в List Redis'a по ключу key
func saveToDataBase(key string, value string) {
	dbClientConn, err := db.Get()
	defer db.Put(dbClientConn)
	if err != nil {
		log.Panic(err)
	}
	dbClientConn.Cmd("RPUSH", key, value)
}
