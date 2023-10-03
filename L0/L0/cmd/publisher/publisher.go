package main

import (
	"L0/internal/config"
	"L0/internal/schema"

	"strconv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {
	cfg := config.InitConfig()

	sc, err := stan.Connect(cfg.NatsCluster, cfg.NatsPublisher, stan.NatsURL(fmt.Sprintf("localhost:%v", cfg.NatsPort)))
	if err != nil {
		log.Fatal("Can't connect to nats", err)
	}
	defer sc.Close()

	for i := 0; i < 10; i++ {
		order := initOrder()

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
			return
		}

		err = sc.Publish(cfg.NatsSubject, bytes)
		if err != nil {
			log.Println(err)
			return
		} else {
			log.Println("Publish order with uid: ", order.OrderUid)
		}
	}
}

func initOrder() *schema.Order {
	id := uuid.New()
	amount, _ := strconv.Atoi(RandStringBytes(4, digits))
	payment, _ := strconv.Atoi(RandStringBytes(4, digits))
	cost, _ := strconv.Atoi(RandStringBytes(5, digits))
	total, _ := strconv.Atoi(RandStringBytes(4, digits))
	fee, _ := strconv.Atoi(RandStringBytes(4, digits))
	chrt, _ := strconv.Atoi(RandStringBytes(10, digits))
	price, _ := strconv.Atoi(RandStringBytes(5, digits))
	sale, _ := strconv.Atoi(RandStringBytes(4, digits))
	nm, _ := strconv.Atoi(RandStringBytes(10, digits))
	status, _ := strconv.Atoi(RandStringBytes(3, digits))
	sm, _ := strconv.Atoi(RandStringBytes(3, digits))

	order := schema.Order{
		OrderUid:    id.String(),
		TrackNumber: RandStringBytes(14, letters),
		Entry:       RandStringBytes(4, letters),
		Delivery: schema.Delivery{
			Name:    RandStringBytes(20, letters),
			Phone:   RandStringBytes(11, digits),
			Zip:     RandStringBytes(7, digits),
			City:    RandStringBytes(20, letters),
			Address: RandStringBytes(20, letters),
			Region:  RandStringBytes(20, letters),
			Email:   RandStringBytes(20, letters),
		},
		Payment: schema.Payment{
			Transaction:  id.String(),
			RequestID:    "",
			Currency:     RandStringBytes(3, letters),
			Provider:     RandStringBytes(5, letters),
			Amount:       amount,
			PaymentDt:    payment,
			Bank:         RandStringBytes(15, letters),
			DeliveryCost: cost,
			GoodsTotal:   total,
			CustomFee:    fee,
		},
		Items: []schema.Item{
			{
				ChrtId:      chrt,
				TrackNumber: RandStringBytes(14, letters),
				Price:       price,
				Rid:         RandStringBytes(21, letters),
				Name:        RandStringBytes(10, letters),
				Sale:        sale,
				Size:        RandStringBytes(1, digits),
				TotalPrice:  total,
				NmId:        nm,
				Brand:       RandStringBytes(10, letters),
				Status:      status,
			},
		},
		Locale:            RandStringBytes(2, letters),
		InternalSignature: RandStringBytes(2, letters),
		CustomerId:        RandStringBytes(10, letters),
		DeliveryService:   RandStringBytes(10, letters),
		Shardkey:          RandStringBytes(2, letters),
		SmId:              sm,
		DateCreated:       time.Now(),
		OofShard:          RandStringBytes(1, letters),
	}

	return &order
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"

func RandStringBytes(n int, digOrLet string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digOrLet[rand.Intn(len(digOrLet))]
	}
	return string(b)
}
