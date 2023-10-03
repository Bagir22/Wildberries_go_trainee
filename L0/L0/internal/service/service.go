package service

import (
	"L0/internal/schema"
	"L0/internal/cache"

	"context"
	"log"
)

type OrderStorage interface {
	Insert(ctx context.Context, order schema.Order, id string) error
	GetOrderById(ctx context.Context, id string) (schema.Order, error)
}

type OrderService struct {
	repo OrderStorage
	cache *cache.InMemory
}

func InitOrderService(repo OrderStorage, cache *cache.InMemory) *OrderService {
	return &OrderService{
		repo: repo,
		cache: cache,
	}
}

func (o *OrderService) Insert(ctx context.Context, order schema.Order, id string) error {
	o.cache.AddToCache(&order)
	return o.repo.Insert(ctx, order, id)
}

func (o *OrderService) GetOrderById(ctx context.Context, id string) (schema.Order, error) {
	data, err := o.cache.GetOrderFromCache(id)
	if err != nil {
		log.Println("Can't get order from cache")
	}

	if data != nil {
		return *data, nil
	}

	return o.repo.GetOrderById(ctx, id)
}
