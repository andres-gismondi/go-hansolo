package services

import (
	log "github.com/sirupsen/logrus"
	"go-hansolo/pkg/server/model"
	"math"
)

type LocationImpl struct {

}

//Epsilon en UNO por las diferencias decimales
const EPSILON = 1

//Coordenadas fijas dadas por el problema
var kenoby = model.Satellite{
	Position: model.Coordinates{X: -500, Y: -200},
}
var Skywalker = model.Satellite{
	Position: model.Coordinates{X: 100, Y: -100},
}
var Sato = model.Satellite{
	Position: model.Coordinates{X: 500, Y: 100},
}

func (LocationImpl) GetLocation(distances ...float64) (x, y float64) {
	/*
	* At this point i dont know wich distances corresponds to each satellite, so
	* i suppose that the first distance corresponds to kenoby and the last to Sato
	 */
	// x0 sky; x1 kenoby; x2 sato
	locationLogger := log.WithFields(log.Fields{"Distances": distances})

	kenobyDistance := distances[0]
	skywalkerDistance := distances[1]
	satoDistance := distances[2]
	dx := kenoby.Position.X - Skywalker.Position.X
	dy := kenoby.Position.Y - Skywalker.Position.Y

	d := math.Sqrt((dy * dy) + (dx * dx))
	if d > (kenobyDistance + skywalkerDistance) {
		locationLogger.Error("Distances between satellites are longer than radius")
		return 0,0
	}
	if d < math.Abs(kenobyDistance-skywalkerDistance) {
		locationLogger.Error("Distances not belong to correct point in common")
		return 0,0
	}

	a := ((kenobyDistance*kenobyDistance) - (skywalkerDistance*skywalkerDistance) + (d*d)) / (2.0 * d)

	//Hipotenusa y demas. Ver bien
	//http://paulbourke.net/geometry/circlesphere/ || fin for "two circles"
	p2x := kenoby.Position.X + (dx * (a/d))
	p2y := kenoby.Position.Y + (dy * (a/d))
	h := math.Sqrt((kenobyDistance*kenobyDistance) - (a*a))
	rx := -dy * (h/d)
	ry := dx * (h/d)

	intersectionP1x := p2x + rx
	intersectionP2x := p2x - rx
	intersectionP1y := p2y + ry
	intersectionP2y := p2y - ry

	//Ahora calcula la distancia del punto de intx con el punto sato
	//si cohinciden alguno de los dos es que hay interseccion entre los tres satellites
	dx = intersectionP1x - Sato.Position.X
	dy = intersectionP1y - Sato.Position.Y
	d1 := math.Sqrt((dy*dy) + (dx*dx))

	dx = intersectionP2x - Sato.Position.X
	dy = intersectionP2y - Sato.Position.Y
	d2 := math.Sqrt((dy*dy) + (dx*dx))

	var resultx float64
	var resulty float64
	if math.Abs(d1-satoDistance) < EPSILON {
		resultx = intersectionP1x
		resulty = intersectionP1y
	}else if math.Abs(d2-satoDistance) < EPSILON {
		resultx = intersectionP2x
		resulty = intersectionP2y
	}else{
		locationLogger.WithField("y", resulty).WithField("x", resultx)
		locationLogger.Error("Can get coordinates")
		return 0, 0
	}

	locationLogger.WithField("y", resulty).WithField("x", resultx).Info("Coordinates founded")
	return resultx, resulty
}
