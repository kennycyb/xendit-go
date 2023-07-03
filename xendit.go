// Package xendit provides the binding for Xendit APIs.
package xendit

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

func NewGenericOpenAPIError(body []byte, error string, model interface{}) *GenericOpenAPIError {
	return &GenericOpenAPIError{
		body: body,
		error: error,
		model: model,
	}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}