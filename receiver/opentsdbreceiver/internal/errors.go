package internal

type JsonDeserializationUnrecoverableError struct {
	e error
}

func (j *JsonDeserializationUnrecoverableError) Error() string {
	return j.e.Error()
}