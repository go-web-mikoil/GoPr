package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	ctx, cansel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cansel()

	return nil
}
