package main

import (
	"fmt"
	"go-hansolo/pkg/server/services"
)

func main() {
	//log.Fatal(http.ListenAndServe(":8080", nil))
	resultX, resultY := services.GetLocation(500, 200, 1538)
	fmt.Println(resultX, resultY)
	m1 := []string{"", "es","","mensaje"}
	m2 := []string{"", "","un",""}
	m3 := []string{"este", "es","",""}
	message := services.GetMessage(m1, m2, m3)
	fmt.Println(message)
}