package webutils

type Event string

func (event Event) String() string {
	return string(event)
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
