package main

import "fmt"

type MemoryKVStore struct {
	store map[string]string
}

type MemStoreOptions struct {
}

func NewMemoryKVStore(options MemStoreOptions) (*MemoryKVStore, error) {
	store := make(map[string]string)
	return &MemoryKVStore{store: store}, nil
}

func (m *MemoryKVStore) Get(key string) (string, error) {
	val, ok := m.store[key]
	if !ok {
		return "", NoSuchKey
	}
	return val, nil
}
func (m *MemoryKVStore) Set(key string, value string) error {
	m.store[key] = value
	return nil
}

func (m *MemoryKVStore) Delete(key string) error {
	delete(m.store, key)
	return nil
}

type MemoryStoreDriver struct {
}

func (d MemoryStoreDriver) NewFromMap(options map[string]interface{}) (KVStore, error) {
	//Nothing to do here
	opts := MemStoreOptions{}
	return NewMemoryKVStore(opts)
}
func (d MemoryStoreDriver) NewFromInterface(options interface{}) (KVStore, error) {
	opts, ok := options.(MemStoreOptions)
	if !ok {
		return nil, fmt.Errorf("Invalid memory options %q", options)
	}
	return NewMemoryKVStore(opts)
}

func init() {
	Register("memory", MemoryStoreDriver{})
}
