package foreverdb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

type Serializer struct {
	maxmb int64
	count int64
	fd    *os.File
	enc   *json.Encoder
	dec   *json.Decoder
	sync.RWMutex
}

func NewSerializer(maxmb int64) *Serializer {
	s := &Serializer{
		maxmb: maxmb,
	}
	s.fd = open(s.count)
	s.enc = json.NewEncoder(s.fd)
	return s
}

func (s *Serializer) Encode(v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.checkUpdateEncoder()
	if err := s.enc.Encode(&v); err != nil {
		log.Fatal("Serializer.Encode(): ", err)
	}
}

func (s *Serializer) Decode() {
	s.Lock()
	defer s.Unlock()
	info, err := ioutil.ReadDir("db")
	if err != nil {
		log.Fatal("Serializer.Decode(): ", err)
	}
	for n, _ := range info {
		fd := open(int64(n))
		defer fd.Close()
		r := bufio.NewReader(fd)
		for {
			line, err := r.ReadBytes('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal("Serializer.Decode(): ", err)
			}
			doc := line[:len(line)-1]
			fmt.Printf("Doc: %s\n", doc)
		}
	}
}

func (s *Serializer) Close() {
	if err := s.fd.Close(); err != nil {
		log.Fatal("Serializer.checkUpdateEncoder() (s.fd.Close()): ", err)
	}
}

func (s *Serializer) checkUpdateEncoder() {
	info, err := s.fd.Stat()
	if err != nil {
		log.Fatal("Serializer.checkUpdateEncoder() (s.fd.Stat()...): ", err)
	}
	if info.Size() < s.maxmb {
		return
	}
	s.count++
	s.fd = open(s.count)
	s.enc = json.NewEncoder(s.fd)
}

func open(fileno int64) *os.File {
	filepath := fmt.Sprintf("db/%d.dat", fileno)
	dir, file := path.Split(filepath)
	if dir != "" {
		if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatal("Serializer.open(): ", err)
			}
		}
	}
	if file != "" {
		if _, err := os.Stat(filepath); err != nil && os.IsNotExist(err) {
			if _, err := os.Create(filepath); err != nil {
				log.Fatal("Serializer.open(): ", err)
			}
		}
	}
	fd, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Serializer.open(): ", err)
	}
	if err := fd.Sync(); err != nil {
		log.Fatal("Serializer.open(): ", err)
	}
	return fd
}
