package main

import (
	"bytes"
	"os"
	"testing"
)

func storeTestHelper(t *testing.T, store KVStore) {
	if _, err := store.Get("test"); err != NoSuchKey {
		t.Fatalf("Expected NoSuchKey, got %q", err)
	}
	if err := store.Set("test", "success"); err != nil {
		t.Fatalf("Expected nil, got %q", err)
	}
	if val, _ := store.Get("test"); val != "success" {
		t.Fatalf("Expected success, got %q", val)
	}
	if err := store.Delete("test"); err != nil {
		t.Fatalf("Expected nil, got %q", err)
	}
	_, err := store.Get("test")
	if err != NoSuchKey {
		t.Fatalf("Expected NoSuchKey, got %q", err)
	}
}

var testMemConfig = []byte(`
{
    "store": {
		"backend": "memory"
	}
}
`)

var testBoltConfig = []byte(`
{
    "store": {
		"backend": "boltdb",
		"options": {
			"filename": "config_test_tempory.db"
		}
	}
}
`)

func TestKVStore(t *testing.T) {
	t.Run("backend=MemoryStore", func(t *testing.T) {
		kv, err := NewStoreFromConfig(bytes.NewReader(testMemConfig))
		if err != nil {
			t.Fatal(err)
		}
		storeTestHelper(t, kv)
	})
	t.Run("backend=BoltDBStore", func(t *testing.T) {
		os.RemoveAll("config_test_tempory.db")       // clean up
		defer os.RemoveAll("config_test_tempory.db") // clean up
		kv, err := NewStoreFromConfig(bytes.NewReader(testBoltConfig))
		if err != nil {
			t.Fatal(err)
		}
		storeTestHelper(t, kv)
	})
}
