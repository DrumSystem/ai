package main

import (
	"sync"
)

type Cache struct {

	rightCount int
	totalCount int
	mu *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		rightCount: 0,
		totalCount: 0,
		mu: &sync.RWMutex{},
	}
}

func (c *Cache) Incr(flag bool)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	if flag {
		c.rightCount++
	}
	c.totalCount++
}

func  (c *Cache) State() (int, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.rightCount, c.totalCount
}

func (c *Cache) Reset() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	r := c.rightCount
	t := c.totalCount
	c.rightCount = 0
	c.totalCount = 0
	return r, t
}