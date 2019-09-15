package calendar

import (
	"errors"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(y int) error {
	if y < 1 {
		return errors.New("invalid year")
	}
	d.Year = y
	return nil
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("invalid month")
	}
	d.Month = m
	return nil
}

func (d *Date) SetDay(a int) error {
	if a < 1 || a > 31 {
		return errors.New("invalid day")
	}
	d.Day = a
	return nil
}
