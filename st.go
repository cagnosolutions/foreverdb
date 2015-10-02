package foreverdb

import "sync"

const mb = 1 << 20

type Store struct {
	name string
	io   *Serializer
	sync.RWMutex
}

func NewStore(name string) *Store {
	return &Store{
		name: name,
		io:   NewSerializer(4 * mb),
	}
}

func (st *Store) Add(v interface{}) {

	return nil
}

func (st *Store) Get(key uint64, v interface{}) error {
	return nil
}

func (st *Store) GetAll(v interface{}) {

}

func (st *Store) Del(key uint64) {

}

func (st *Store) AddBin(v []byte) {

}

func (st *Store) GetBin(key uint64) ([]byte, error) {
	return nil, nil
}

func (st *Store) GetAllBin() [][]byte {
	return nil
}
