// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// UserDeleteOKCode is the HTTP code returned for type UserDeleteOK
const UserDeleteOKCode int = 200

/*UserDeleteOK item is deleted

swagger:response userDeleteOK
*/
type UserDeleteOK struct {
}

// NewUserDeleteOK creates UserDeleteOK with default headers values
func NewUserDeleteOK() *UserDeleteOK {

	return &UserDeleteOK{}
}

// WriteResponse to the client
func (o *UserDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UserDeleteBadRequestCode is the HTTP code returned for type UserDeleteBadRequest
const UserDeleteBadRequestCode int = 400

/*UserDeleteBadRequest invalid input, object invalid

swagger:response userDeleteBadRequest
*/
type UserDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserDeleteBadRequest creates UserDeleteBadRequest with default headers values
func NewUserDeleteBadRequest() *UserDeleteBadRequest {

	return &UserDeleteBadRequest{}
}

// WithPayload adds the payload to the user delete bad request response
func (o *UserDeleteBadRequest) WithPayload(payload *models.Error) *UserDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user delete bad request response
func (o *UserDeleteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserDeleteUnauthorizedCode is the HTTP code returned for type UserDeleteUnauthorized
const UserDeleteUnauthorizedCode int = 401

/*UserDeleteUnauthorized bad authentication

swagger:response userDeleteUnauthorized
*/
type UserDeleteUnauthorized struct {
}

// NewUserDeleteUnauthorized creates UserDeleteUnauthorized with default headers values
func NewUserDeleteUnauthorized() *UserDeleteUnauthorized {

	return &UserDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *UserDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UserDeleteForbiddenCode is the HTTP code returned for type UserDeleteForbidden
const UserDeleteForbiddenCode int = 403

/*UserDeleteForbidden bad permissions

swagger:response userDeleteForbidden
*/
type UserDeleteForbidden struct {
}

// NewUserDeleteForbidden creates UserDeleteForbidden with default headers values
func NewUserDeleteForbidden() *UserDeleteForbidden {

	return &UserDeleteForbidden{}
}

// WriteResponse to the client
func (o *UserDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UserDeleteConflictCode is the HTTP code returned for type UserDeleteConflict
const UserDeleteConflictCode int = 409

/*UserDeleteConflict item already exists

swagger:response userDeleteConflict
*/
type UserDeleteConflict struct {
}

// NewUserDeleteConflict creates UserDeleteConflict with default headers values
func NewUserDeleteConflict() *UserDeleteConflict {

	return &UserDeleteConflict{}
}

// WriteResponse to the client
func (o *UserDeleteConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// UserDeleteServiceUnavailableCode is the HTTP code returned for type UserDeleteServiceUnavailable
const UserDeleteServiceUnavailableCode int = 503

/*UserDeleteServiceUnavailable internal server error

swagger:response userDeleteServiceUnavailable
*/
type UserDeleteServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserDeleteServiceUnavailable creates UserDeleteServiceUnavailable with default headers values
func NewUserDeleteServiceUnavailable() *UserDeleteServiceUnavailable {

	return &UserDeleteServiceUnavailable{}
}

// WithPayload adds the payload to the user delete service unavailable response
func (o *UserDeleteServiceUnavailable) WithPayload(payload *models.Error) *UserDeleteServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user delete service unavailable response
func (o *UserDeleteServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserDeleteServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
