package main

import (
	"sync"
)

func computingAndSaving(in string) {
	var commonMap = make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		firstOperation(in, commonMap)
	}()
	wg.Wait()
	startSymbolsString := getStartSymbols(in)
	for i := int64(0); i < repeatInt; i++ {
		go insertKey(commonMap, startSymbolsString)
	}
}

func firstOperation(in string, inMap map[string]string) {
	hash := genHash(in)
	json := buildJsonResponse(in, hash)
	inMap[getFourLastSymbols(in)] = hash
	saving(sourceValueString, json)
}