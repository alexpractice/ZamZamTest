package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

//Модель конфиг-файла
type Config struct {
	WsPort   string         `json:"wsPort"`
	DbConfig DataBaseConfig `json:"dataBase"`
}

//Модель конфигурации БД в конфиг-файле
type DataBaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

//Загрузка конфиг-файла по его относительному пути
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
