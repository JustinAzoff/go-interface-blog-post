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

type StoreConfig struct {
	Backend string                 `json:"backend"`
	Options map[string]interface{} `json:"options"`
}

type Configuration struct {
	Store StoreConfig `json:"store"`
}

func NewStore(backend string, options interface{}) (KVStore, error) {
	switch backend {
	case "memory":
		opts, ok := options.(MemStoreOptions)
		if !ok {
			return nil, fmt.Errorf("Invalid memory options %q", options)
		}
		return NewMemoryKVStore(opts)
	case "boltdb":
		opts, ok := options.(BoltStoreOptions)
		if !ok {
			return nil, fmt.Errorf("Invalid Bolt options %q", options)
		}
		return NewBoltStore(opts)
	default:
		return nil, fmt.Errorf("Unknown store: %s", backend)
	}
}

func NewStoreFromConfig(r io.Reader) (KVStore, error) {
	var cfg Configuration
	err := json.NewDecoder(r).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	switch cfg.Store.Backend {
	case "memory":
		return NewMemoryStoreFromMap(cfg.Store.Options)
	case "boltdb":
		return NewBoltStoreFromMap(cfg.Store.Options)
	default:
		return nil, fmt.Errorf("Unknown store: %s", cfg.Store.Backend)
	}
}
