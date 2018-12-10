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

//функция генерации случайного числа (не больше 4 знаков)
func genRand() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Int() % 10000
	return r1
}

//добавлению в мапу уже сгенерированных "окончаний" уникального ключа
func insertKey(set map[string]string, start string) {
	var insertFlag = false
	for insertFlag != true {
		a := genRand()
		s := strconv.FormatInt(int64(a), 10)
		s = verifyStringNumber(s) //ЧИСЛО ДОЛЖНО ИМЕТЬ 4 ЗНАКА
		lock.Lock()
		defer lock.Unlock()
		if value, ok := set[s]; ok {
			log.Println(value, "is exists in map")
		} else {
			fin := start + s
			hash := genHash(fin)
			set[s] = hash
			json := buildJsonResponse(fin, hash)
			saving(json)
			insertFlag = true
		}
	}
}

//функция делает число 4-ёх разрядным, если сгенерированное число имеет меньше 4-ёх символов
func verifyStringNumber(in string) string {
	inputLength := 4 - len(in)
	for inputLength > 0 {
		in = "0" + in
		inputLength--
	}
	return in
}
