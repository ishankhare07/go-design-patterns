package builder

import (
	"testing"
)

func TestBuildingCar(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	manufacturingComplex.SetBuilder(&CarBuilder{})
	car := manufacturingComplex.Construct()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4. Got=%d", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Car must have a 'Car' structure. Got = %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Car must have 5 seats. Got = %d", car.Seats)
	}
}

func TestBuildingBike(t *testing.T) {
	manufacturingComplex := &ManufacturingDirector{}

	manufacturingComplex.SetBuilder(&BikeBuilder{})
	bike := manufacturingComplex.Construct()

	if bike.Wheels != 2 {
		t.Errorf("Wheels in a bike must be 2. Got = %d", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Bike must have a 'Bike' structure. Got = %s", bike.Structure)
	}

	if bike.Seats != 2 {
		t.Errorf("Bike must have 2 seats. Got = %d", bike.Seats)
	}
}
