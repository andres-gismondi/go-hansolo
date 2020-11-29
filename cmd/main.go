package main

import (
	"go-hansolo/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	//log.Fatal(http.ListenAndServe(":8080", nil))
	/*resultX, resultY := services.GetLocation(500, 200, 1538)
	fmt.Println(resultX, resultY)
	m1 := []string{"", "es","","mensaje"}
	m2 := []string{"", "","un",""}
	m3 := []string{"este", "es","",""}
	message := services.GetMessage(m1, m2, m3)
	fmt.Println(message)*/

	//port := os.Getenv("PORT")
	serv, err := server.New("8083")
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