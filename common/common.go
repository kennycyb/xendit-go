package common

import (
	"context"
	"net/http"
	"net/url"
)

type IConfiguration interface {
	AddDefaultHeader(key string, value string)
	ServerURL(index int, variables map[string]string) (string, error)
	ServerURLWithContext(ctx context.Context, endpoint string) (string, error)
}

type IClient interface {
	GetConfig() IConfiguration
	PrepareRequest(
		ctx context.Context,
		path string, method string,
		postBody interface{},
		headerParams map[string]string,
		queryParams url.Values,
		formParams url.Values,
		formFiles []FormFile) (localVarRequest *http.Request, err error)
	CallAPI(request *http.Request) (*http.Response, error)
	Decode(v interface{}, b []byte, contentType string) (err error)
}

type FormFile struct {
	FileBytes []byte
	FileName string
	FormFileName string
}

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