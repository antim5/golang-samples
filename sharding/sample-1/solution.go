package main

import (
	"fmt"
	"sync"
)

type Cache[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (value V, ok bool)
}

var _ Cache[int, int] = (*cacheImpl[int, int])(nil)

type cacheImpl[K comparable, V any] struct {
	data  map[K]V
	mutex sync.Mutex
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return &cacheImpl[K, V]{
		data: map[K]V{},
	}
}

func (impl *cacheImpl[K, V]) Set(key K, value V) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()
	impl.data[key] = value
}

func (impl *cacheImpl[K, V]) Get(key K) (value V, ok bool) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()
	value, ok = impl.data[key]
	return
}

func main() {
	cahce := NewCache[string, int]()
	cahce.Set("a", 1)
	cahce.Set("b", 2)
	v, ok := cahce.Get("a")
	fmt.Printf("cahce[%s] = %d %v\n", "a", v, ok)
}
