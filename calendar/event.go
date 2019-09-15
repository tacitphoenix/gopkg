package calendar

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(t string) {
	e.title = t
}

func (e *Event) Title() string {
	return e.title
}
