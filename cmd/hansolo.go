package cmd

import (
	"go-hansolo/internal/server"
	"log"
)

func Hansolo() {
	port := "8080"
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}
	serv.Start()
}