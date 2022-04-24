package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufactuingComplex := ManufacturingDirector{}

	CarBuilder := &CarBuilder{}

	manufactuingComplex.SetBuilder(CarBuilder)
	manufactuingComplex.Construct()

	car := CarBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	// Test for motorbike

	bikeBuilder := &BikeBuilder{}
	manufactuingComplex.SetBuilder(bikeBuilder)
	manufactuingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on motorbike must be 2 and they were %d", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s", motorbike.Structure)
	}

}
