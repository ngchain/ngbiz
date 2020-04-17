// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetTxpoolCheckHashOKCode is the HTTP code returned for type GetTxpoolCheckHashOK
const GetTxpoolCheckHashOKCode int = 200

/*GetTxpoolCheckHashOK OK

swagger:response getTxpoolCheckHashOK
*/
type GetTxpoolCheckHashOK struct {
}

// NewGetTxpoolCheckHashOK creates GetTxpoolCheckHashOK with default headers values
func NewGetTxpoolCheckHashOK() *GetTxpoolCheckHashOK {

	return &GetTxpoolCheckHashOK{}
}

// WriteResponse to the client
func (o *GetTxpoolCheckHashOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetTxpoolCheckHashBadRequestCode is the HTTP code returned for type GetTxpoolCheckHashBadRequest
const GetTxpoolCheckHashBadRequestCode int = 400

/*GetTxpoolCheckHashBadRequest Error

swagger:response getTxpoolCheckHashBadRequest
*/
type GetTxpoolCheckHashBadRequest struct {

	/*error text
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetTxpoolCheckHashBadRequest creates GetTxpoolCheckHashBadRequest with default headers values
func NewGetTxpoolCheckHashBadRequest() *GetTxpoolCheckHashBadRequest {

	return &GetTxpoolCheckHashBadRequest{}
}

// WithPayload adds the payload to the get txpool check hash bad request response
func (o *GetTxpoolCheckHashBadRequest) WithPayload(payload string) *GetTxpoolCheckHashBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get txpool check hash bad request response
func (o *GetTxpoolCheckHashBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTxpoolCheckHashBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
