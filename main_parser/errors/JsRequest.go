package errors

type JsRequestError struct {
	Message string
}

func New_js_request_error(text string) error {
	return &JsRequestError{text}
}

func (e *JsRequestError) Error() string {
	return e.Message
}
