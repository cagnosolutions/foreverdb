package foreverdb

import (
	"errors"
	"sync"
)

var (
	ErrStoreNotExist = errors.New("store does not exist")
	ErrStoreExists   = errors.New("store already exists")
)

type DataStore struct {
	stores map[string]*Store
	sync.RWMutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		stores: make(map[string]*Store),
	}
}

func (ds *DataStore) AddStore(name string) error {
	st := ds.getstore(name)
	if st != nil {
		return ErrStoreExists
	}
	ds.Lock()
	defer ds.Unlock()
	ds.stores[name] = NewStore(name)
	return nil
}

func (ds *DataStore) GetStore(name string) (*Store, error) {
	st := ds.getstore(name)
	if st == nil {
		return nil, ErrStoreNotExist
	}
	return st, nil
}

func (ds *DataStore) DelStore(name string) {
	st := ds.getstore(name)
	if st != nil {
		ds.Lock()
		st = nil // GC
		delete(ds.stores, name)
		ds.Unlock()
	}
}

func (ds *DataStore) getstore(name string) *Store {
	ds.RLock()
	defer ds.RUnlock()
	if st, ok := ds.stores[name]; ok {
		return st
	}
	return nil
}
