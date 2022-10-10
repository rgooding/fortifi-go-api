// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetProducts(params *GetProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOK, error)

	GetProductsGroups(params *GetProductsGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsGroupsOK, error)

	GetProductsGroupsProductGroupFidProducts(params *GetProductsGroupsProductGroupFidProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsGroupsProductGroupFidProductsOK, error)

	GetProductsOffers(params *GetProductsOffersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOffersOK, error)

	GetProductsProductFidPricebands(params *GetProductsProductFidPricebandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidPricebandsOK, error)

	GetProductsProductFidPrices(params *GetProductsProductFidPricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidPricesOK, error)

	GetProductsProductFidSkus(params *GetProductsProductFidSkusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidSkusOK, error)

	PostProductsProductFidAvailabilityCheck(params *PostProductsProductFidAvailabilityCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProductsProductFidAvailabilityCheckOK, error)

	PostProductsProductFidAvailabilityReserve(params *PostProductsProductFidAvailabilityReserveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProductsProductFidAvailabilityReserveOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetProducts lists all products
*/
func (a *Client) GetProducts(params *GetProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsReader{formats: a.formats},
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
	success, ok := result.(*GetProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsGroups gets a list of all product groups
*/
func (a *Client) GetProductsGroups(params *GetProductsGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsGroups",
		Method:             "GET",
		PathPattern:        "/products/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetProductsGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsGroupsProductGroupFidProducts gets a list of products belonging to the group
*/
func (a *Client) GetProductsGroupsProductGroupFidProducts(params *GetProductsGroupsProductGroupFidProductsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsGroupsProductGroupFidProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsGroupsProductGroupFidProductsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsGroupsProductGroupFidProducts",
		Method:             "GET",
		PathPattern:        "/products/groups/{productGroupFid}/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsGroupsProductGroupFidProductsReader{formats: a.formats},
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
	success, ok := result.(*GetProductsGroupsProductGroupFidProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsGroupsProductGroupFidProductsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsOffers retrieves all offers
*/
func (a *Client) GetProductsOffers(params *GetProductsOffersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsOffersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsOffersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsOffers",
		Method:             "GET",
		PathPattern:        "/products/offers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsOffersReader{formats: a.formats},
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
	success, ok := result.(*GetProductsOffersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsOffersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsProductFidPricebands retrieves product price bands
*/
func (a *Client) GetProductsProductFidPricebands(params *GetProductsProductFidPricebandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidPricebandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsProductFidPricebandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsProductFidPricebands",
		Method:             "GET",
		PathPattern:        "/products/{productFid}/pricebands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsProductFidPricebandsReader{formats: a.formats},
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
	success, ok := result.(*GetProductsProductFidPricebandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsProductFidPricebandsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsProductFidPrices retrieves product prices
*/
func (a *Client) GetProductsProductFidPrices(params *GetProductsProductFidPricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsProductFidPricesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsProductFidPrices",
		Method:             "GET",
		PathPattern:        "/products/{productFid}/prices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsProductFidPricesReader{formats: a.formats},
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
	success, ok := result.(*GetProductsProductFidPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsProductFidPricesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProductsProductFidSkus retrieves product s k us
*/
func (a *Client) GetProductsProductFidSkus(params *GetProductsProductFidSkusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProductsProductFidSkusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsProductFidSkusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProductsProductFidSkus",
		Method:             "GET",
		PathPattern:        "/products/{productFid}/skus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductsProductFidSkusReader{formats: a.formats},
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
	success, ok := result.(*GetProductsProductFidSkusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProductsProductFidSkusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostProductsProductFidAvailabilityCheck checks product availability
*/
func (a *Client) PostProductsProductFidAvailabilityCheck(params *PostProductsProductFidAvailabilityCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProductsProductFidAvailabilityCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProductsProductFidAvailabilityCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostProductsProductFidAvailabilityCheck",
		Method:             "POST",
		PathPattern:        "/products/{productFid}/availability/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostProductsProductFidAvailabilityCheckReader{formats: a.formats},
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
	success, ok := result.(*PostProductsProductFidAvailabilityCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostProductsProductFidAvailabilityCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostProductsProductFidAvailabilityReserve reserves product
*/
func (a *Client) PostProductsProductFidAvailabilityReserve(params *PostProductsProductFidAvailabilityReserveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProductsProductFidAvailabilityReserveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProductsProductFidAvailabilityReserveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostProductsProductFidAvailabilityReserve",
		Method:             "POST",
		PathPattern:        "/products/{productFid}/availability/reserve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostProductsProductFidAvailabilityReserveReader{formats: a.formats},
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
	success, ok := result.(*PostProductsProductFidAvailabilityReserveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostProductsProductFidAvailabilityReserveDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
