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

var locationLogger log.Logger

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

	dx, dy, d := GetDistances()
	if d > (kenobyDistance + skywalkerDistance) {
		locationLogger.Error("Distances between satellites are longer than radius")
		return 0,0
	}
	if d < math.Abs(kenobyDistance-skywalkerDistance) {
		locationLogger.Error("Distances not belong to correct point in common")
		return 0,0
	}

	a := ((kenobyDistance*kenobyDistance) - (skywalkerDistance*skywalkerDistance) + (d*d)) / (2.0 * d)
	intersectionP1x, intersectionP2x, intersectionP1y, intersectionP2y := GetIntersectionPoints(dx, dy, a, d, kenobyDistance)

	//Ahora calcula la distancia del punto de intx con el punto sato
	//si cohinciden alguno de los dos es que hay interseccion entre los tres satellites
	d1, d2 := GetLastPointIntersection(dx, dy, intersectionP1x, intersectionP1y, intersectionP2x, intersectionP2y)

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

func GetDistances() (dx, dy , d float64) {
	dx = kenoby.Position.X - Skywalker.Position.X
	dy = kenoby.Position.Y - Skywalker.Position.Y

	d = math.Sqrt((dy * dy) + (dx * dx))
	return dx, dy, d
}

func GetIntersectionPoints(dx, dy, a, d, kDistance float64) (p1x, p1y, p2x, p2y float64) {
	//Hipotenusa y demas. Ver bien
	//http://paulbourke.net/geometry/circlesphere/ || fin for "two circles"
	p2x = kenoby.Position.X + (dx * (a/d))
	p2y = kenoby.Position.Y + (dy * (a/d))
	h := math.Sqrt((kDistance * kDistance) - (a*a))
	rx := -dy * (h/d)
	ry := dx * (h/d)
	return p2x + rx, p2x - rx, p2y + ry, p2y - ry
}

func GetLastPointIntersection(dx, dy, intP1x, intP1y, intP2x, intP2y float64) (d1, d2 float64) {
	dx = intP1x - Sato.Position.X
	dy = intP1y - Sato.Position.Y
	d1 = math.Sqrt((dy*dy) + (dx*dx))

	dx = intP2x - Sato.Position.X
	dy = intP2y - Sato.Position.Y
	d2 = math.Sqrt((dy*dy) + (dx*dx))

	return d1, d2
}
