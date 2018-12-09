package main

import "sync"

func saving(src string, json string) {
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

func saveToChannel(json string) {
	channel <- json
}