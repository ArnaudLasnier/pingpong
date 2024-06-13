package webutils

type Event string

func (event Event) String() string {
	return string(event)
}
