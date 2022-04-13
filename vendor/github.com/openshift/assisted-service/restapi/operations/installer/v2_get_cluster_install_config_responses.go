// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetClusterInstallConfigOKCode is the HTTP code returned for type V2GetClusterInstallConfigOK
const V2GetClusterInstallConfigOKCode int = 200

/*V2GetClusterInstallConfigOK Success.

swagger:response v2GetClusterInstallConfigOK
*/
type V2GetClusterInstallConfigOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigOK creates V2GetClusterInstallConfigOK with default headers values
func NewV2GetClusterInstallConfigOK() *V2GetClusterInstallConfigOK {

	return &V2GetClusterInstallConfigOK{}
}

// WithPayload adds the payload to the v2 get cluster install config o k response
func (o *V2GetClusterInstallConfigOK) WithPayload(payload string) *V2GetClusterInstallConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config o k response
func (o *V2GetClusterInstallConfigOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2GetClusterInstallConfigUnauthorizedCode is the HTTP code returned for type V2GetClusterInstallConfigUnauthorized
const V2GetClusterInstallConfigUnauthorizedCode int = 401

/*V2GetClusterInstallConfigUnauthorized Unauthorized.

swagger:response v2GetClusterInstallConfigUnauthorized
*/
type V2GetClusterInstallConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigUnauthorized creates V2GetClusterInstallConfigUnauthorized with default headers values
func NewV2GetClusterInstallConfigUnauthorized() *V2GetClusterInstallConfigUnauthorized {

	return &V2GetClusterInstallConfigUnauthorized{}
}

// WithPayload adds the payload to the v2 get cluster install config unauthorized response
func (o *V2GetClusterInstallConfigUnauthorized) WithPayload(payload *models.InfraError) *V2GetClusterInstallConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config unauthorized response
func (o *V2GetClusterInstallConfigUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetClusterInstallConfigForbiddenCode is the HTTP code returned for type V2GetClusterInstallConfigForbidden
const V2GetClusterInstallConfigForbiddenCode int = 403

/*V2GetClusterInstallConfigForbidden Forbidden.

swagger:response v2GetClusterInstallConfigForbidden
*/
type V2GetClusterInstallConfigForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigForbidden creates V2GetClusterInstallConfigForbidden with default headers values
func NewV2GetClusterInstallConfigForbidden() *V2GetClusterInstallConfigForbidden {

	return &V2GetClusterInstallConfigForbidden{}
}

// WithPayload adds the payload to the v2 get cluster install config forbidden response
func (o *V2GetClusterInstallConfigForbidden) WithPayload(payload *models.InfraError) *V2GetClusterInstallConfigForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config forbidden response
func (o *V2GetClusterInstallConfigForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetClusterInstallConfigNotFoundCode is the HTTP code returned for type V2GetClusterInstallConfigNotFound
const V2GetClusterInstallConfigNotFoundCode int = 404

/*V2GetClusterInstallConfigNotFound Error.

swagger:response v2GetClusterInstallConfigNotFound
*/
type V2GetClusterInstallConfigNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigNotFound creates V2GetClusterInstallConfigNotFound with default headers values
func NewV2GetClusterInstallConfigNotFound() *V2GetClusterInstallConfigNotFound {

	return &V2GetClusterInstallConfigNotFound{}
}

// WithPayload adds the payload to the v2 get cluster install config not found response
func (o *V2GetClusterInstallConfigNotFound) WithPayload(payload *models.Error) *V2GetClusterInstallConfigNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config not found response
func (o *V2GetClusterInstallConfigNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetClusterInstallConfigMethodNotAllowedCode is the HTTP code returned for type V2GetClusterInstallConfigMethodNotAllowed
const V2GetClusterInstallConfigMethodNotAllowedCode int = 405

/*V2GetClusterInstallConfigMethodNotAllowed Method Not Allowed.

swagger:response v2GetClusterInstallConfigMethodNotAllowed
*/
type V2GetClusterInstallConfigMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigMethodNotAllowed creates V2GetClusterInstallConfigMethodNotAllowed with default headers values
func NewV2GetClusterInstallConfigMethodNotAllowed() *V2GetClusterInstallConfigMethodNotAllowed {

	return &V2GetClusterInstallConfigMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 get cluster install config method not allowed response
func (o *V2GetClusterInstallConfigMethodNotAllowed) WithPayload(payload *models.Error) *V2GetClusterInstallConfigMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config method not allowed response
func (o *V2GetClusterInstallConfigMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetClusterInstallConfigInternalServerErrorCode is the HTTP code returned for type V2GetClusterInstallConfigInternalServerError
const V2GetClusterInstallConfigInternalServerErrorCode int = 500

/*V2GetClusterInstallConfigInternalServerError Error.

swagger:response v2GetClusterInstallConfigInternalServerError
*/
type V2GetClusterInstallConfigInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetClusterInstallConfigInternalServerError creates V2GetClusterInstallConfigInternalServerError with default headers values
func NewV2GetClusterInstallConfigInternalServerError() *V2GetClusterInstallConfigInternalServerError {

	return &V2GetClusterInstallConfigInternalServerError{}
}

// WithPayload adds the payload to the v2 get cluster install config internal server error response
func (o *V2GetClusterInstallConfigInternalServerError) WithPayload(payload *models.Error) *V2GetClusterInstallConfigInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get cluster install config internal server error response
func (o *V2GetClusterInstallConfigInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetClusterInstallConfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
