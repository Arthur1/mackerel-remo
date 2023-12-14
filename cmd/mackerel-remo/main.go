package main

import (
	"context"
	"log"

	mackerelremo "github.com/Arthur1/mackerel-remo"
	"github.com/caarlos0/env"
)

func main() {
	cfg := &mackerelremo.RunnerConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	runner := mackerelremo.NewRunner(cfg)
	ctx := context.Background()
	if err := runner.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
