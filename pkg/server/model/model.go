package model

type RequestModel struct {
	Satellites []RequestSatellite `json:"satellites"`
}

type RequestSatellite struct {
	Name string			`json:"name"`
	Distance float64	`json:"distance"`
	Message []string	`json:"message"`
}

type ResponseModel struct {
	Position Coordinates	`json:"position"`
	Message string			`json:"message"`
}

type Coordinates struct {
	X float64	`json:"x"`
	Y float64	`json:"y"`
}

type Satellite struct {
	Position Coordinates
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Map map[string]interface{}