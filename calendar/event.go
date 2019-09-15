package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(t string) error {
	if utf8.RuneCountInString(t) > 30 {
		return errors.New("invalid title because it's too long")
	}
	e.title = t
	return nil
}

func (e *Event) Title() string {
	return e.title
}
