package geo

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) SetLatitude(val float64) {
	c.latitude = val
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) SetLongitude(val float64) {
	c.longitude = val
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}
