package main

import (
	"encoding/json"
	"log"
)

type jsonScheme struct {
	Number string
	Hash   string
}

//функция создания json'ов согласно jsonScheme
func buildJsonResponse(num string, hash string) string {
	myScheme := &jsonScheme{
		Number: num,
		Hash:   hash}
	realJson, err := json.Marshal(myScheme)
	if err != nil {
		log.Fatal(err)
	}
	return string(realJson)
}
