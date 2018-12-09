// Code generated by go-swagger; DO NOT EDIT.

package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stevexnicholls/next/models"
)

// ValueGetOKCode is the HTTP code returned for type ValueGetOK
const ValueGetOKCode int = 200

/*ValueGetOK successful operation

swagger:response valueGetOK
*/
type ValueGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.KeyValue `json:"body,omitempty"`
}

// NewValueGetOK creates ValueGetOK with default headers values
func NewValueGetOK() *ValueGetOK {

	return &ValueGetOK{}
}

// WithPayload adds the payload to the value get o k response
func (o *ValueGetOK) WithPayload(payload *models.KeyValue) *ValueGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value get o k response
func (o *ValueGetOK) SetPayload(payload *models.KeyValue) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ValueGetNotFoundCode is the HTTP code returned for type ValueGetNotFound
const ValueGetNotFoundCode int = 404

/*ValueGetNotFound The entry was not found

swagger:response valueGetNotFound
*/
type ValueGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewValueGetNotFound creates ValueGetNotFound with default headers values
func NewValueGetNotFound() *ValueGetNotFound {

	return &ValueGetNotFound{}
}

// WithPayload adds the payload to the value get not found response
func (o *ValueGetNotFound) WithPayload(payload *models.Error) *ValueGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value get not found response
func (o *ValueGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ValueGetDefault Error

swagger:response valueGetDefault
*/
type ValueGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewValueGetDefault creates ValueGetDefault with default headers values
func NewValueGetDefault(code int) *ValueGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ValueGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the value get default response
func (o *ValueGetDefault) WithStatusCode(code int) *ValueGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the value get default response
func (o *ValueGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the value get default response
func (o *ValueGetDefault) WithPayload(payload *models.Error) *ValueGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value get default response
func (o *ValueGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
