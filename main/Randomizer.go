package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	lock = sync.Mutex{}
)

func genRand() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Int() % 10000
	return r1
}

func insertKey(set map[string]string, start string) {
	var insertFlag = false;
	for insertFlag != true {
		a := genRand()
		s := strconv.FormatInt(int64(a), 10)
		lock.Lock()
		defer lock.Unlock()
		if value, ok := set[s]; ok {
			log.Println(value, "is exists in map")
		} else {
			s = verifyStringNumber(s) //ЧИСЛО ДОЛЖНО ИМЕТЬ 4 ЗНАКА
			fin := start + s
			hash := genHash(fin)
			set[s] = hash
			json := buildJsonResponse(fin, hash)
			saving(sourceValueString, json)
			insertFlag = true
		}
	}
}

func verifyStringNumber(in string) string {
	inputLength := 4 - len(in)
	for ; inputLength > 0; {
		in = "0" + in
		inputLength--
	}
	return in
}
