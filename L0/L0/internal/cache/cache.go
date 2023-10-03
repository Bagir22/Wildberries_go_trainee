package cache

import (
	"L0/internal/schema"

	"errors"
	"sync"
)

type InMemory struct {
	sync.RWMutex
	Cache map[string]*schema.Order
}

func InitCache() *InMemory {
	c := &InMemory{}
	c.Cache = make(map[string]*schema.Order)

	return c
}

func (c *InMemory) AddToCache(order *schema.Order) error {
	c.Lock()
	defer c.Unlock()

	c.Cache[order.OrderUid] = order

	return nil
}

func (c *InMemory) GetOrderFromCache(id string) (*schema.Order, error) {
	c.RLock()
	order, ok := c.Cache[id]
	defer c.RUnlock()

	if !ok {
		return nil, errors.New("Missing order in cache")
	}

	return order, nil
}
