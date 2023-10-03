package postgres

import (
	"L0/internal/schema"
	"L0/internal/cache"
	"L0/internal/config"
	
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func InitDb(db *sqlx.DB) *Db {
	return &Db{
		db: db,
	}
}

func InitConn() (*sqlx.DB, error) {
	cfg := config.InitConfig()

	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDatabase, "disable")

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Db) Insert(ctx context.Context, order schema.Order, orderUid string) error {

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Println("Can't marshal order")
		return err
	}

	_, err = d.db.Exec("INSERT INTO orders (data) values ($1);", jsonOrder)
	if err != nil {
		log.Println("Cant't insert order into database")
		return err
	}
	return nil
}

func (d *Db) GetOrderById(ctx context.Context, orderUid string) (schema.Order, error) {
	var order schema.Order
	row := d.db.QueryRow("SELECT data FROM orders WHERE order_uid = $1", orderUid)
	err := row.Scan(&order)

	if err != nil {
		log.Println("Can't get order from database")
		return schema.Order{}, err
	}

	return order, nil
}

func (d *Db) RecoverCache(ctx context.Context, cache *cache.InMemory) error {
	rows, err := d.db.Query("SELECT data FROM orders")
	if err != nil {
		log.Println("Can't recover cache")
		return err
	}
	
	for rows.Next() {
		var order schema.Order
		err := rows.Scan(&order)
		if err != nil {
			log.Println("Can't scan values to Order struct")
			return err
		}

		if err = cache.AddToCache(&order); err != nil {
			log.Println("Can't add order to cache")
			return err
		}
	}

	return nil
}
