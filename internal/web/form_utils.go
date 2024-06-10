package web

type Form struct {
	IsSubmitted bool
	Fields      FormFields
}

type FormFields map[FormFieldKey]FormFieldValue

type FormFieldValue struct {
	Value   string
	IsValid bool
	Message string
}

func NewValidValue(value string) FormFieldValue {
	return FormFieldValue{
		Value:   value,
		IsValid: true,
		Message: FormFieldOKMessage,
	}
}

func NewInvalidValue(value string, message string) FormFieldValue {
	return FormFieldValue{
		Value:   value,
		IsValid: false,
		Message: message,
	}
}

const FormFieldOKMessage = "Looks good!"

type FormFieldKey string

func (key FormFieldKey) String() string {
	return string(key)
}

const (
	PlayerFirstName FormFieldKey = "firstName"
	PlayerLastName  FormFieldKey = "lastName"
	PlayerEmail     FormFieldKey = "email"
	TournamentTitle FormFieldKey = "title"
)
