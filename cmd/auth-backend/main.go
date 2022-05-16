package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Stop to gracefully shutdown application upon signal
func Stop() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	done := <-shutdown
	fmt.Printf("Received shutdown signal %s\n", done.String())
	os.Exit(0)
}

func main() {
	logger := log.Default()
	handler, err := InitializeAndRun(context.Background(), logger, "config.env")
	if err != nil {
		logger.Println(context.Background(), "Failed to initialize the app", "err", err)
		os.Exit(2)
	}

	go http.ListenAndServe(":3000", handler.Router)
	fmt.Println("Listening on localhost:3000...")
	Stop()
}
