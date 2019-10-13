package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
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
