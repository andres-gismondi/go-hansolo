package services

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"strings"
)

type MessageImpl struct {

}

func (MessageImpl) GetMessage(messages ...[]string) (string, error) {
	messagesLogger := log.WithFields(log.Fields{"messages": messages})
	sentence, condition := ConcatenateSlices(messages)
	if !condition {
		messagesLogger.Error("can not decode message")
		return "can not decode message", errors.New("can not decode message")
	}
	msg := strings.Join(sentence, " ")
	messagesLogger.Info("message got it")
	return msg, nil
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
