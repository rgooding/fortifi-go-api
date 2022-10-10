// Code generated by go-swagger; DO NOT EDIT.

package reviews

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reviews API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reviews API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PostReview(params *PostReviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostReviewOK, error)

	PutReview(params *PutReviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutReviewOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PostReview creates a new review and review audit
*/
func (a *Client) PostReview(params *PostReviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostReviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostReview",
		Method:             "POST",
		PathPattern:        "/review",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostReviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostReviewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutReview updates a review
*/
func (a *Client) PutReview(params *PutReviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutReviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutReviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutReview",
		Method:             "PUT",
		PathPattern:        "/review",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutReviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutReviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutReviewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}