package cache

import (
	"sync"

	"github.com/jackyzha0/ctrl-v/db"
)

type Cache struct {
	m    map[string]db.Paste
	lock sync.RWMutex
}

func New() *Cache {
	return &Cache{
		m: map[string]db.Paste{},
	}
}

func (c *Cache) Get(hash string) (db.Paste, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// check if hash in cache
	if v, ok := c.m[hash]; ok {
		return v, nil
	}

	// if it doesnt, lookup from db
	p, err := db.Lookup(hash)
	c.add(p)
	return p, err
}

func (c *Cache) add(p db.Paste) {
	c.lock.Lock()
	c.m[p.Hash] = p
	c.lock.Unlock()
}