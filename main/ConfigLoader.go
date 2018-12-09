package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	WsPort   string         `json:"wsPort"`
	DbConfig DataBaseConfig `json:"dataBase"`
}

type DataBaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(fileRelativePath string) Config {
	absPath, err := filepath.Abs(fileRelativePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	var config Config
	configFile, err := os.Open(absPath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config

}
