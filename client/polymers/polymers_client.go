// Code generated by go-swagger; DO NOT EDIT.

package polymers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new polymers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for polymers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPolymersPolymerFid(params *GetPolymersPolymerFidParams, authInfo runtime.ClientAuthInfoWriter) (*GetPolymersPolymerFidOK, error)

	PostPolymers(params *PostPolymersParams, authInfo runtime.ClientAuthInfoWriter) (*PostPolymersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPolymersPolymerFid reads a polymer
*/
func (a *Client) GetPolymersPolymerFid(params *GetPolymersPolymerFidParams, authInfo runtime.ClientAuthInfoWriter) (*GetPolymersPolymerFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolymersPolymerFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolymersPolymerFid",
		Method:             "GET",
		PathPattern:        "/polymers/{polymerFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolymersPolymerFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPolymersPolymerFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPolymersPolymerFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostPolymers creates a new polymer
*/
func (a *Client) PostPolymers(params *PostPolymersParams, authInfo runtime.ClientAuthInfoWriter) (*PostPolymersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPolymersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPolymers",
		Method:             "POST",
		PathPattern:        "/polymers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPolymersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostPolymersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostPolymersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
