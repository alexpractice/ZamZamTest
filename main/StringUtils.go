package main

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"strconv"
)

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

func getFourLastSymbols(s string) string {
	length := len(s)
	substring := s[length-4 : length]
	return substring
}

func getStartSymbols(s string) string {
	length := len(s)
	substring := s[:length-4]
	return substring
}
