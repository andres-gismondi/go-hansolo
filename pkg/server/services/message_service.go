package services

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type MessageImpl struct {

}

func (MessageImpl) GetMessage(messages ...[]string) (mes string) {
	messagesLogger := log.WithFields(log.Fields{"Messages": messages})
	sentence, condition := ConcatenateSlices(messages)
	if !condition {
		messagesLogger.Error("Can not decode message")
		return "Can not decode message"
	}
	msg := strings.Join(sentence, " ")
	messagesLogger.Info("Message got it")
	return msg
}

func ConcatenateSlices(messages [][]string) ([]string, bool) {
	var cleanSlice []string

	anotherSlice := InitializeSlice(len(messages[0]))
	for _, str := range messages {
		cleanSlice = append(cleanSlice, str...)
		for i := range(str) {
			if str[i] != "" {
				anotherSlice[i] = str[i]
			}
		}
	}
	return anotherSlice, FindEmptyItem(anotherSlice)
}

func InitializeSlice(length int)  []string{
	conditions := make([]string, length)
	for i:=0; i<length; i++ {
		conditions[i] = ""
	}
	return conditions
}

func FindEmptyItem(slice []string) bool {
	for _, item := range slice {
		if item == "" {
			return false
		}
	}
	return true
}
