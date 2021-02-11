package absfactory

type LuxuaryCar struct{}

func (l *LuxuaryCar) NumDoors() int {
	return 4
}

func (l *LuxuaryCar) NumWheels() int {
	return 4
}

func (l *LuxuaryCar) NumSeats() int {
	return 5
}
