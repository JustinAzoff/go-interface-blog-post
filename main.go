package main

import "fmt"

func main() {
	kv, _ := NewMemoryKVStore(MemStoreOptions{})

	fmt.Println(kv.Get("test"))
	fmt.Println(kv.Set("test", "success"))
	fmt.Println(kv.Get("test"))
	fmt.Println(kv.Delete("test"))
	fmt.Println(kv.Get("test"))
}
