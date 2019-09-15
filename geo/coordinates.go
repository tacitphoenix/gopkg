package geo

import "errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) SetLatitude(val float64) error {
	if val < -90 || val > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = val
	return nil
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) SetLongitude(val float64) error {
	if val < -180 || val > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = val
	return nil
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}
