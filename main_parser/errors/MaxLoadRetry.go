package errors

type MaxLoadRetryError struct {
	Message string
}

func New_max_load_retry_error(text string) error {
	return &MaxLoadRetryError{text}
}

func (e *MaxLoadRetryError) Error() string {
	return e.Message
}
