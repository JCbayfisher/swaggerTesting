// Code generated by go-swagger; DO NOT EDIT.

package logout

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/JCbayfisher/swaggerTesting/models"
)

// PostLogoutOKCode is the HTTP code returned for type PostLogoutOK
const PostLogoutOKCode int = 200

/*PostLogoutOK Logout successful

swagger:response postLogoutOK
*/
type PostLogoutOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostLogoutOK creates PostLogoutOK with default headers values
func NewPostLogoutOK() *PostLogoutOK {

	return &PostLogoutOK{}
}

// WithPayload adds the payload to the post logout o k response
func (o *PostLogoutOK) WithPayload(payload *models.ReturnCode) *PostLogoutOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post logout o k response
func (o *PostLogoutOK) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLogoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLogoutBadRequestCode is the HTTP code returned for type PostLogoutBadRequest
const PostLogoutBadRequestCode int = 400

/*PostLogoutBadRequest Bad Request

swagger:response postLogoutBadRequest
*/
type PostLogoutBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostLogoutBadRequest creates PostLogoutBadRequest with default headers values
func NewPostLogoutBadRequest() *PostLogoutBadRequest {

	return &PostLogoutBadRequest{}
}

// WithPayload adds the payload to the post logout bad request response
func (o *PostLogoutBadRequest) WithPayload(payload *models.ReturnCode) *PostLogoutBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post logout bad request response
func (o *PostLogoutBadRequest) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLogoutBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLogoutNotFoundCode is the HTTP code returned for type PostLogoutNotFound
const PostLogoutNotFoundCode int = 404

/*PostLogoutNotFound Not Found

swagger:response postLogoutNotFound
*/
type PostLogoutNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostLogoutNotFound creates PostLogoutNotFound with default headers values
func NewPostLogoutNotFound() *PostLogoutNotFound {

	return &PostLogoutNotFound{}
}

// WithPayload adds the payload to the post logout not found response
func (o *PostLogoutNotFound) WithPayload(payload *models.ReturnCode) *PostLogoutNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post logout not found response
func (o *PostLogoutNotFound) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLogoutNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLogoutInternalServerErrorCode is the HTTP code returned for type PostLogoutInternalServerError
const PostLogoutInternalServerErrorCode int = 500

/*PostLogoutInternalServerError Internal Server Error

swagger:response postLogoutInternalServerError
*/
type PostLogoutInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostLogoutInternalServerError creates PostLogoutInternalServerError with default headers values
func NewPostLogoutInternalServerError() *PostLogoutInternalServerError {

	return &PostLogoutInternalServerError{}
}

// WithPayload adds the payload to the post logout internal server error response
func (o *PostLogoutInternalServerError) WithPayload(payload *models.ReturnCode) *PostLogoutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post logout internal server error response
func (o *PostLogoutInternalServerError) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLogoutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostLogoutDefault Logout unexpected error response

swagger:response postLogoutDefault
*/
type PostLogoutDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostLogoutDefault creates PostLogoutDefault with default headers values
func NewPostLogoutDefault(code int) *PostLogoutDefault {
	if code <= 0 {
		code = 500
	}

	return &PostLogoutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post logout default response
func (o *PostLogoutDefault) WithStatusCode(code int) *PostLogoutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post logout default response
func (o *PostLogoutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post logout default response
func (o *PostLogoutDefault) WithPayload(payload *models.ReturnCode) *PostLogoutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post logout default response
func (o *PostLogoutDefault) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLogoutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
