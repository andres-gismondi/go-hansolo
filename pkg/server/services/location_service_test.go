package services

import "testing"

func TestGetLocation(t *testing.T) {
	service := &LocationImpl{}
	distances := []float64{500,500,500}

	//Cant get coordinates
	_, _, err := service.GetLocation(distances...)
	if err == nil {
		t.Fatal("Should be an error")
	}

	//Success
	distances = []float64{500,200,1538}
	x, y, err := service.GetLocation(distances...)
	if err != nil {
		t.Fatal("Should not be an error")
	}
	if x == 0 && y == 0{
		t.Fatal("X and y should not be zero")
	}
}
