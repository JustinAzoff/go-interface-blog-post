package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type KVStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

var NoSuchKey = errors.New("No such key")

type KVStoreDriver interface {
	NewFromInterface(interface{}) (KVStore, error)
	NewFromMap(map[string]interface{}) (KVStore, error)
}

type StoreConfig struct {
	Backend string                 `json:"backend"`
	Options map[string]interface{} `json:"options"`
}

type Configuration struct {
	Store StoreConfig `json:"store"`
}

type StoreBackend struct {
	options MemStoreOptions
}

var drivers = make(map[string]KVStoreDriver)

func Register(backend string, driver KVStoreDriver) {
	drivers[backend] = driver
}

func NewStore(backend string, options interface{}) (KVStore, error) {
	driver, ok := drivers[backend]
	if !ok {
		return nil, fmt.Errorf("Unknown store: %s", backend)
	}
	return driver.NewFromInterface(options)
}

func NewStoreFromConfig(r io.Reader) (KVStore, error) {
	var cfg Configuration
	err := json.NewDecoder(r).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	driver, ok := drivers[cfg.Store.Backend]
	if !ok {
		return nil, fmt.Errorf("Unknown store: %s", cfg.Store.Backend)
	}
	return driver.NewFromMap(cfg.Store.Options)
}
