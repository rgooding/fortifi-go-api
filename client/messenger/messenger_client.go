// Code generated by go-swagger; DO NOT EDIT.

package messenger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new messenger API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for messenger API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PutMessengerDeliveriesDeliveryFidUnsubscribe(params *PutMessengerDeliveriesDeliveryFidUnsubscribeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutMessengerDeliveriesDeliveryFidUnsubscribeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PutMessengerDeliveriesDeliveryFidUnsubscribe unsubscribes an email based on the delivery fid
*/
func (a *Client) PutMessengerDeliveriesDeliveryFidUnsubscribe(params *PutMessengerDeliveriesDeliveryFidUnsubscribeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutMessengerDeliveriesDeliveryFidUnsubscribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMessengerDeliveriesDeliveryFidUnsubscribeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutMessengerDeliveriesDeliveryFidUnsubscribe",
		Method:             "PUT",
		PathPattern:        "/messenger/deliveries/{deliveryFid}/unsubscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutMessengerDeliveriesDeliveryFidUnsubscribeReader{formats: a.formats},
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
	success, ok := result.(*PutMessengerDeliveriesDeliveryFidUnsubscribeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutMessengerDeliveriesDeliveryFidUnsubscribeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
