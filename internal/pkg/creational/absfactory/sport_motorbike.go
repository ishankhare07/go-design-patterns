package absfactory

type SportsMotorbike struct{}

func (s *SportsMotorbike) GetMotorbikeType() int {
	return SportsMotorbikeType
}

func (s *SportsMotorbike) NumWheels() int {
	return 2
}

func (s *SportsMotorbike) NumSeats() int {
	return 1
}
