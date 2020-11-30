package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage(t *testing.T) {
	service := &MessageImpl{}

	//Success
	kenobiMessage := []string{"este","","","mensaje",""}
	skywalkerMessage := []string{"","es","","",""}
	satoMessage := []string{"","","un","","secreto"}

	messages := [][]string{
		kenobiMessage,
		skywalkerMessage,
		satoMessage,
	}
	msg, err := service.GetMessage(messages...)
	if err != nil {
		t.Fatal("An error occurred")
	}
	assert.Equal(t, "este es un mensaje secreto", msg)
}

func TestErrorMessage(t *testing.T) {
	service := &MessageImpl{}

	//A word is missing (secreto)

	kenobiMessage := []string{"este","","","mensaje",""}
	skywalkerMessage := []string{"","es","","",""}
	satoMessage := []string{"","","un","",""}

	messages := [][]string{
		kenobiMessage,
		skywalkerMessage,
		satoMessage,
	}
	msg, err := service.GetMessage(messages...)
	if err == nil {
		t.Fatal("An error occurred")
	}
	assert.Equal(t, "can not decode message", msg)
}
