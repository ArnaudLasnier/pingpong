package webutils

import "strings"

type Event string

func (event Event) String() string {
	return string(event)
}

func (event Event) FromBody() string {
	return event.String() + " from:body"
}

func JoinEvents(events ...Event) string {
	var eventStrs []string
	for _, event := range events {
		eventStrs = append(eventStrs, event.FromBody())
	}
	return strings.Join(eventStrs, ",")
}

type SuccessData struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type ErrorData struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
