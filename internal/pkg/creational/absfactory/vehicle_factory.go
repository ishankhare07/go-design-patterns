package absfactory

import "fmt"

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

type CarFactory struct{}

const (
	LuxuryCarType = iota
	FamilyCarType
)

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuaryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

const (
	SportsMotorbikeType = iota
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportsMotorbikeType:
		return new(SportsMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

const (
	CarFactoryType = iota
	MotorbikeFactoryType
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("Factory with id %d not recognized\n", f)
	}
}
