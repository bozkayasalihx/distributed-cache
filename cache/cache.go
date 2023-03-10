package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	Cacher
	data map[string][]byte
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[string(key)]
	if !ok {
		return []byte(""), nil
	}
	return val, nil
}

func (c *Cache) Set(key []byte, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[string(key)] = val
	return nil

}

func (c *Cache) Update(key []byte, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[string(key)]
	if !ok {
		return fmt.Errorf("key not found %s", string(key))
	}

	c.data[string(key)] = val
	return nil
}

func (c *Cache) Delete(key []byte) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[string(key)]
	if !ok {
		return []byte(""), fmt.Errorf("couldn't find the %s", string(key))
	}
	return val, nil

}
