package builder

// BuildProcess defines all the
// steps required to be implemented
// by a builder
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector takes a BuildProcess
// and constructs a VehicleProduct using that builder
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct constructs a VehicleProduct
// out of a BuildProcess
func (m *ManufacturingDirector) Construct() VehicleProduct {
	return m.builder.
		SetWheels().
		SetSeats().
		SetStructure().
		GetVehicle()
}

// SetBuilder sets the builder to be used for
// building the VehicleProduct
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

// VehicleProduct is the actual Vehicle object
// returned by all builders
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder implements the BuildProcess
// interface for Cars
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels sets the wheels for the car
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats sets the number of seats for the car
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure sets the structure for car
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle returns the underlying VehicleProduct
// for the car
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// BikeBuilder implemets the BuildProcess
// interface for a Bike
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels sets the number of wheels for bike
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats sets number of seats for a bike
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure sets the structure for a bike
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

// GetVehicle returns a VehicleProduct for a bike
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// BusBuilder implements the BuildProcess
// interface for a bus
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels sets the number of wheels on the bus
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 6
	return b
}

// SetSeats sets the number of seats on the bus
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 40
	return b
}

// SetStructure sets the structure of the bus
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

// GetVehicle returns the underlying VehicleProduct
// for the bus
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
