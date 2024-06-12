package web

type Form struct {
	IsSubmitted bool
	Fields      FormFields
}

type FormFields map[FormKey]FormValue

type FormKey string

func (key FormKey) String() string {
	return string(key)
}

type FormValue struct {
	Value             string
	IsValid           bool
	ValidationMessage string
}

const validationMessageOK = "Looks good!"

func NewValidFormValue(value string) FormValue {
	return FormValue{
		Value:             value,
		IsValid:           true,
		ValidationMessage: validationMessageOK,
	}
}

func NewInvalidFormValue(value string, message string) FormValue {
	return FormValue{
		Value:             value,
		IsValid:           false,
		ValidationMessage: message,
	}
}
