package calendar

import (
	"errors"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(y int) error {
	if y < 1 {
		return errors.New("invalid year")
	}
	d.year = y
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("invalid month")
	}
	d.month = m
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetDay(a int) error {
	if a < 1 || a > 31 {
		return errors.New("invalid day")
	}
	d.day = a
	return nil
}

func (d *Date) Day() int {
	return d.day
}
