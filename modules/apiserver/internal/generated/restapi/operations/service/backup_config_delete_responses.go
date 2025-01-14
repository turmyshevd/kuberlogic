// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// BackupConfigDeleteOKCode is the HTTP code returned for type BackupConfigDeleteOK
const BackupConfigDeleteOKCode int = 200

/*BackupConfigDeleteOK item is deleted

swagger:response backupConfigDeleteOK
*/
type BackupConfigDeleteOK struct {
}

// NewBackupConfigDeleteOK creates BackupConfigDeleteOK with default headers values
func NewBackupConfigDeleteOK() *BackupConfigDeleteOK {

	return &BackupConfigDeleteOK{}
}

// WriteResponse to the client
func (o *BackupConfigDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// BackupConfigDeleteBadRequestCode is the HTTP code returned for type BackupConfigDeleteBadRequest
const BackupConfigDeleteBadRequestCode int = 400

/*BackupConfigDeleteBadRequest invalid input, object invalid

swagger:response backupConfigDeleteBadRequest
*/
type BackupConfigDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigDeleteBadRequest creates BackupConfigDeleteBadRequest with default headers values
func NewBackupConfigDeleteBadRequest() *BackupConfigDeleteBadRequest {

	return &BackupConfigDeleteBadRequest{}
}

// WithPayload adds the payload to the backup config delete bad request response
func (o *BackupConfigDeleteBadRequest) WithPayload(payload *models.Error) *BackupConfigDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config delete bad request response
func (o *BackupConfigDeleteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupConfigDeleteUnauthorizedCode is the HTTP code returned for type BackupConfigDeleteUnauthorized
const BackupConfigDeleteUnauthorizedCode int = 401

/*BackupConfigDeleteUnauthorized bad authentication

swagger:response backupConfigDeleteUnauthorized
*/
type BackupConfigDeleteUnauthorized struct {
}

// NewBackupConfigDeleteUnauthorized creates BackupConfigDeleteUnauthorized with default headers values
func NewBackupConfigDeleteUnauthorized() *BackupConfigDeleteUnauthorized {

	return &BackupConfigDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *BackupConfigDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BackupConfigDeleteForbiddenCode is the HTTP code returned for type BackupConfigDeleteForbidden
const BackupConfigDeleteForbiddenCode int = 403

/*BackupConfigDeleteForbidden bad permissions

swagger:response backupConfigDeleteForbidden
*/
type BackupConfigDeleteForbidden struct {
}

// NewBackupConfigDeleteForbidden creates BackupConfigDeleteForbidden with default headers values
func NewBackupConfigDeleteForbidden() *BackupConfigDeleteForbidden {

	return &BackupConfigDeleteForbidden{}
}

// WriteResponse to the client
func (o *BackupConfigDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// BackupConfigDeleteNotFoundCode is the HTTP code returned for type BackupConfigDeleteNotFound
const BackupConfigDeleteNotFoundCode int = 404

/*BackupConfigDeleteNotFound item not found

swagger:response backupConfigDeleteNotFound
*/
type BackupConfigDeleteNotFound struct {
}

// NewBackupConfigDeleteNotFound creates BackupConfigDeleteNotFound with default headers values
func NewBackupConfigDeleteNotFound() *BackupConfigDeleteNotFound {

	return &BackupConfigDeleteNotFound{}
}

// WriteResponse to the client
func (o *BackupConfigDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// BackupConfigDeleteServiceUnavailableCode is the HTTP code returned for type BackupConfigDeleteServiceUnavailable
const BackupConfigDeleteServiceUnavailableCode int = 503

/*BackupConfigDeleteServiceUnavailable internal server error

swagger:response backupConfigDeleteServiceUnavailable
*/
type BackupConfigDeleteServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigDeleteServiceUnavailable creates BackupConfigDeleteServiceUnavailable with default headers values
func NewBackupConfigDeleteServiceUnavailable() *BackupConfigDeleteServiceUnavailable {

	return &BackupConfigDeleteServiceUnavailable{}
}

// WithPayload adds the payload to the backup config delete service unavailable response
func (o *BackupConfigDeleteServiceUnavailable) WithPayload(payload *models.Error) *BackupConfigDeleteServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config delete service unavailable response
func (o *BackupConfigDeleteServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigDeleteServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
