package main

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"strconv"
)

//функция генерирования хэша
func genHash(input string) string {
	hash := sha3.NewKeccak256()

	var buf []byte
	b, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	hash.Write(b)
	buf = hash.Sum(buf)

	return hex.EncodeToString(buf)
}

//фунция валидации числа, переданного в качестве первого аргумента
func verifyInputInt(in string) {
	a, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}
	//if a%2 != 0 {
	//	panic(errors.New("can't work with odd integer"))
	//}
	if a/100000 < 1 {
		panic(errors.New("can't work with integer, which have less 6 digits"))
	}
}

//функция получения последних четырёх цифр числа для использования в качестве ключа в map'е
func getFourLastSymbols(s string) string {
	length := len(s)
	substring := s[length-4 : length]
	return substring
}

//функция получения первых цифр исходного числа, к которым будут добавляться сгенерированные четрыре
func getStartSymbols(s string) string {
	length := len(s)
	substring := s[:length-4]
	return substring
}
