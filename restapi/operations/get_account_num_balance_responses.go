// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAccountNumBalanceOKCode is the HTTP code returned for type GetAccountNumBalanceOK
const GetAccountNumBalanceOKCode int = 200

/*GetAccountNumBalanceOK OK

swagger:response getAccountNumBalanceOK
*/
type GetAccountNumBalanceOK struct {
}

// NewGetAccountNumBalanceOK creates GetAccountNumBalanceOK with default headers values
func NewGetAccountNumBalanceOK() *GetAccountNumBalanceOK {

	return &GetAccountNumBalanceOK{}
}

// WriteResponse to the client
func (o *GetAccountNumBalanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetAccountNumBalanceBadRequestCode is the HTTP code returned for type GetAccountNumBalanceBadRequest
const GetAccountNumBalanceBadRequestCode int = 400

/*GetAccountNumBalanceBadRequest Error

swagger:response getAccountNumBalanceBadRequest
*/
type GetAccountNumBalanceBadRequest struct {

	/*error text
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAccountNumBalanceBadRequest creates GetAccountNumBalanceBadRequest with default headers values
func NewGetAccountNumBalanceBadRequest() *GetAccountNumBalanceBadRequest {

	return &GetAccountNumBalanceBadRequest{}
}

// WithPayload adds the payload to the get account num balance bad request response
func (o *GetAccountNumBalanceBadRequest) WithPayload(payload string) *GetAccountNumBalanceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account num balance bad request response
func (o *GetAccountNumBalanceBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountNumBalanceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
