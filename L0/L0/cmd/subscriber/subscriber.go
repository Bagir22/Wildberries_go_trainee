package main

import (
	"L0/internal/cache"
	"L0/internal/config"
	"L0/internal/handler"
	"L0/internal/postgres"
	"L0/internal/schema"
	"L0/internal/server"
	"L0/internal/service"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	cfg := config.InitConfig()

	const timeout = time.Second * 10

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := postgres.InitConn()
	if err != nil {
		log.Fatal("Can't init connection to database")
	}
	defer conn.Close()

	db := postgres.InitDb(conn)

	cache := cache.InitCache()

	db.RecoverCache(ctx, cache)

	service := service.InitOrderService(db, cache)

	handler := handler.InitHandler(service)

	server.Run(handler)

	nc, err := stan.Connect(cfg.NatsCluster, cfg.NatsSubscriber, stan.NatsURL(fmt.Sprintf("localhost:%v", cfg.NatsPort)))

	if err != nil {
		log.Fatal("Can't connect to nats streaming", err)
	}

	_, err = nc.Subscribe(cfg.NatsSubject, func(m *stan.Msg) {
		var order schema.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("Can't unmarshall json")
			return
		}

		err = service.Insert(ctx, order, order.OrderUid)
		if err != nil {
			log.Println(err)
			return
		}
	})

	if err != nil {
		log.Println(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	<-c
}
