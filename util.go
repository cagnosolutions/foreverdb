package foreverdb

import (
	"encoding/binary"
	"encoding/json"
	"log"
)

// possible data structure scheme to use?: http://play.golang.org/p/5_bro0h0Z5
// how to set id for struct automatically: http://play.golang.org/p/mFd-8tVXx6

// also, consider using a node tree (checkout gabs: github.com/Jeffail/gabs)...

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
