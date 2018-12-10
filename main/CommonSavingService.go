package main

import "sync"

//функция сохранения полученного json'а в БД и канал
func saving(json string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		saveToDataBase(sourceValueString, json)
	}()
	go func() {
		defer wg.Done()
		saveToChannel(json)
	}()
	wg.Wait()
}

//передача строчки в канал
func saveToChannel(json string) {
	channel <- json
}
