package main

import (
	"log"
	"time"

	"github.com/geekAshish/go-grpc-graphql-micro/catalog"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(i int) error {
		_, err := catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	defer r.Close()

	log.Println("listing on PORT 8080...")

	s := catalog.NewService(r)

	log.Fatal(catalog.ListenGRPC(s, 8080))
}
