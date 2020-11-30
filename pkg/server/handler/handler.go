package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"go-hansolo/pkg/server/model"
	"go-hansolo/pkg/server/services"
	"net/http"
)

type Impl struct {
	LocationService 	services.LocationImpl
	MessageService 		services.MessageImpl
}

var (
	distances 		[]float64
	messages 		[][]string
	Kenoby 			model.RequestSatellite
	Skywalker 		model.RequestSatellite
	Sato 			model.RequestSatellite
)

func (s *Impl) TopSecretHandler(writer http.ResponseWriter, request *http.Request) {
	log.Info("Getting location and messages")
	var rm model.RequestModel
	err := json.NewDecoder(request.Body).Decode(&rm)

	if err != nil {
		_ = HTTPError(writer, http.StatusNotFound, err.Error())
		return
	}

	distances = []float64{rm.Satellites[0].Distance,rm.Satellites[1].Distance,rm.Satellites[2].Distance}
	messages = [][]string{rm.Satellites[0].Message,rm.Satellites[1].Message,rm.Satellites[2].Message}
	s.GetService(distances,messages,writer)
}

func (s *Impl) GetSplitHandler(writer http.ResponseWriter, request *http.Request) {
	satelliteName := chi.URLParam(request, "satellite_name")
	log.Info("Get split with for satellite: " + satelliteName)

	if satelliteName != "kenoby" && satelliteName != "skywalker" && satelliteName != "sato" {
		_ = HTTPError(writer, http.StatusNotFound, "Satellite name not found")
		return
	}

	kenobyName := &Kenoby.Name
	skywalkerName := &Skywalker.Name
	satoName := &Sato.Name
	if *kenobyName != "" && *skywalkerName != "" && *satoName != "" {
		distances = []float64{Kenoby.Distance, Skywalker.Distance, Sato.Distance}
		messages = [][]string{Kenoby.Message, Skywalker.Message, Sato.Message}
		s.GetService(distances,messages,writer)
	} else {
		_ = HTTPError(writer, http.StatusNotFound, "Not enough information to get data")
		return
	}
}

func (s *Impl) PostSplitHandler(writer http.ResponseWriter, request *http.Request) {
	var rm model.RequestSatellite
	err := json.NewDecoder(request.Body).Decode(&rm)
	if err != nil {
		_ = HTTPError(writer, http.StatusNotFound, err.Error())
		return
	}

	satelliteName := chi.URLParam(request, "satellite_name")
	log.Info("Post satellite: " + satelliteName)
	rm.Name = satelliteName

	if satelliteName == "kenoby" {
		Kenoby.Name = satelliteName
		Kenoby.Distance = rm.Distance
		Kenoby.Message = rm.Message
	} else if satelliteName == "skywalker" {
		Skywalker.Name = satelliteName
		Skywalker.Distance = rm.Distance
		Skywalker.Message = rm.Message
	} else if satelliteName == "sato" {
		Sato.Name = satelliteName
		Sato.Distance = rm.Distance
		Sato.Message = rm.Message
	} else {
		_ = HTTPError(writer, http.StatusNotFound, "Satellite name not found")
		return
	}

	_ = JSON(writer, http.StatusAccepted, "Ok")
}

func (s *Impl) GetService(distances []float64, messages [][]string, writer http.ResponseWriter) {
	//500, 200, 1538
	rX, rY, err := s.LocationService.GetLocation(
		distances[0],
		distances[1],
		distances[2])
	if err != nil {
		_ = HTTPError(writer, http.StatusNotFound, "Cant get coordinates")
		return
	}

	msg, err := s.MessageService.GetMessage(
		messages[0],
		messages[1],
		messages[2])
	if err != nil {
		_ = HTTPError(writer, http.StatusNotFound, "Cant get message")
		return
	}

	response := model.ResponseModel{
		Position: model.Coordinates{X: rX, Y: rY},
		Message: msg,
	}
	_ = JSON(writer, http.StatusAccepted, response)
}

func (s *Impl) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/topsecret", s.TopSecretHandler)
	r.Post("/topsecret_split/{satellite_name}", s.PostSplitHandler)
	r.Get("/topsecret_split/{satellite_name}", s.GetSplitHandler)
	return r
}