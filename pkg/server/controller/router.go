package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go-hansolo/pkg/server/model"
	"go-hansolo/pkg/server/services"
	"net/http"
)

type Impl struct {

}

func (s *Impl) GetLocationHandler(writer http.ResponseWriter, request *http.Request) {
	var rm model.RequestModel
	err := json.NewDecoder(request.Body).Decode(&rm)
	if err != nil {
		_ = HTTPError(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	//600, 520, 658
	//500, 200, 1538
	rX, rY := services.GetLocation(
		rm.Satellites[0].Distance,
		rm.Satellites[1].Distance,
		rm.Satellites[2].Distance)

	msg := services.GetMessage(
		rm.Satellites[0].Message,
		rm.Satellites[1].Message,
		rm.Satellites[2].Message)

	if rX == 0 && rY == 0 {
		_ = HTTPError(writer, request, http.StatusBadRequest, "Cant get coordinates")
		return
	}

	response := model.ResponseModel{
		Position: model.Coordinates{X: rX, Y: rY},
		Message: msg,
	}
	_ = JSON(writer, request, http.StatusAccepted, response)
}

func (s *Impl) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/topsecret", s.GetLocationHandler)
	return r
}