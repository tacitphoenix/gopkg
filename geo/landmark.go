package geo

import "errors"

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) SetName(val string) error {
	if val == "" {
		return errors.New("invalid name can't be blank")
	}
	l.name = val
	return nil
}

func (l *Landmark) Name() string {
	return l.name
}
