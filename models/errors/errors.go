package errors

type ServerError struct {
	msg string
}

func (s *ServerError) Error() string {
	return s.msg
}

func New(msg string) error {

	err := &ServerError{
		msg: msg,
	}

	return err
}
