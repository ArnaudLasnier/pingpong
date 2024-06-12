package web

type form struct {
	IsSubmitted bool
	Fields      formFields
}

type formFields map[formKey]formValue

type formKey string

func (key formKey) String() string {
	return string(key)
}

type formValue struct {
	Value             string
	IsValid           bool
	ValidationMessage string
}

const validationMessageOK = "Looks good!"

func newValidFormValue(value string) formValue {
	return formValue{
		Value:             value,
		IsValid:           true,
		ValidationMessage: validationMessageOK,
	}
}

func newInvalidFormValue(value string, message string) formValue {
	return formValue{
		Value:             value,
		IsValid:           false,
		ValidationMessage: message,
	}
}
