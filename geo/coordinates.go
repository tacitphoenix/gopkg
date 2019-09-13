package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinates) SetLatitude(val float64) {
	c.Latitude = val
}

func (c *Coordinates) SetLongitude(val float64) {
	c.Longitude = val
}
