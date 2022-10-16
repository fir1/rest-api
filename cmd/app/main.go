package main

import (
	"context"
	"github.com/fir1/rest-api/http/rest"
	"log"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := rest.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
