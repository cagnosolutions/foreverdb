package foreverdb

import (
	"encoding/binary"
	"encoding/json"
	"log"
)

func UintToBytes(n uint64) []byte {
	b := make([]byte, 10)
	binary.PutUvarint(b, n)
	return b
}

func EncodeJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func DecodeJSON(b []byte, v interface{}) {
	err := json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}
}
