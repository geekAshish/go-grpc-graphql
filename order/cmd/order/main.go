package main

import (
	"log"
	"time"

	"github.com/geekAshish/go-grpc-graphql-micro/order"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		return
	})
	defer r.Close()

	log.Println("Server is listning on port 8080")
	s := order.NewService(r)
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, 8080))
}
