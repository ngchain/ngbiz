// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*GetOK OK

swagger:response getOK
*/
type GetOK struct {
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}