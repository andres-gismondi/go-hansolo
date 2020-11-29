package main

import (
	"go-hansolo/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {

	port := "8089"
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	_ = serv.Close()
}