package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go-hansolo/pkg/server/model"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) error {
	if data == nil {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(statusCode)
		log.Warn("Empty data")
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		log.Error("Error marshaling data")
		return err
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(statusCode)
	writer.Write(j)
	return nil
}

func HTTPError(writer http.ResponseWriter, statusCode int, message string) error {
	msg := model.ErrorResponse{
		Message: message,
	}
	return JSON(writer, statusCode, msg)
}