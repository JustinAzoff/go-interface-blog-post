package main

type MemoryKVStore struct {
	store map[string]string
}

type MemStoreOptions struct {
}

func NewMemoryKVStore(options MemStoreOptions) (*MemoryKVStore, error) {
	store := make(map[string]string)
	return &MemoryKVStore{store: store}, nil
}

func NewMemoryStoreFromMap(options map[string]interface{}) (*MemoryKVStore, error) {
	//Nothing to do here
	opts := MemStoreOptions{}
	return NewMemoryKVStore(opts)
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
