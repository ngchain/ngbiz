// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostTxpoolSendRawTxOKCode is the HTTP code returned for type PostTxpoolSendRawTxOK
const PostTxpoolSendRawTxOKCode int = 200

/*PostTxpoolSendRawTxOK OK

swagger:response postTxpoolSendRawTxOK
*/
type PostTxpoolSendRawTxOK struct {
}

// NewPostTxpoolSendRawTxOK creates PostTxpoolSendRawTxOK with default headers values
func NewPostTxpoolSendRawTxOK() *PostTxpoolSendRawTxOK {

	return &PostTxpoolSendRawTxOK{}
}

// WriteResponse to the client
func (o *PostTxpoolSendRawTxOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostTxpoolSendRawTxBadRequestCode is the HTTP code returned for type PostTxpoolSendRawTxBadRequest
const PostTxpoolSendRawTxBadRequestCode int = 400

/*PostTxpoolSendRawTxBadRequest Error

swagger:response postTxpoolSendRawTxBadRequest
*/
type PostTxpoolSendRawTxBadRequest struct {
}

// NewPostTxpoolSendRawTxBadRequest creates PostTxpoolSendRawTxBadRequest with default headers values
func NewPostTxpoolSendRawTxBadRequest() *PostTxpoolSendRawTxBadRequest {

	return &PostTxpoolSendRawTxBadRequest{}
}

// WriteResponse to the client
func (o *PostTxpoolSendRawTxBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}