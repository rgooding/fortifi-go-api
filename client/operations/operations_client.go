// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetIntegrationsVerifyUser(params *GetIntegrationsVerifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntegrationsVerifyUserOK, error)

	GetPayCoinbase(params *GetPayCoinbaseParams, authInfo runtime.ClientAuthInfoWriter) (*GetPayCoinbaseOK, error)

	GetPayPublicKey(params *GetPayPublicKeyParams, authInfo runtime.ClientAuthInfoWriter) (*GetPayPublicKeyOK, error)

	PostCustomersCustomerFidAddresses(params *PostCustomersCustomerFidAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidAddressesOK, error)

	PostCustomersCustomerFidEmails(params *PostCustomersCustomerFidEmailsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidEmailsOK, error)

	PostCustomersCustomerFidInvoicesInvoiceFidCreditNote(params *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteOK, error)

	PostCustomersCustomerFidPhones(params *PostCustomersCustomerFidPhonesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidPhonesOK, error)

	PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidEventsOK, error)

	GetBrands(params *GetBrandsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBrandsOK, error)

	GetMe(params *GetMeParams, authInfo runtime.ClientAuthInfoWriter) (*GetMeOK, error)

	GetOrganisation(params *GetOrganisationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganisationOK, error)

	GetServiceAuthToken(params *GetServiceAuthTokenParams) (*GetServiceAuthTokenOK, error)

	GetVersion(params *GetVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetIntegrationsVerifyUser verifies a user
*/
func (a *Client) GetIntegrationsVerifyUser(params *GetIntegrationsVerifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntegrationsVerifyUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntegrationsVerifyUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIntegrationsVerifyUser",
		Method:             "GET",
		PathPattern:        "/integrations/verifyUser",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntegrationsVerifyUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntegrationsVerifyUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIntegrationsVerifyUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPayCoinbase gets a new checkout ID
*/
func (a *Client) GetPayCoinbase(params *GetPayCoinbaseParams, authInfo runtime.ClientAuthInfoWriter) (*GetPayCoinbaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayCoinbaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayCoinbase",
		Method:             "GET",
		PathPattern:        "/pay/coinbase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayCoinbaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPayCoinbaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPayCoinbaseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPayPublicKey gets the public key needed to encrypt a credit card number
*/
func (a *Client) GetPayPublicKey(params *GetPayPublicKeyParams, authInfo runtime.ClientAuthInfoWriter) (*GetPayPublicKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayPublicKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayPublicKey",
		Method:             "GET",
		PathPattern:        "/pay/publicKey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayPublicKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPayPublicKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPayPublicKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCustomersCustomerFidAddresses adds an address to a customer
*/
func (a *Client) PostCustomersCustomerFidAddresses(params *PostCustomersCustomerFidAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidAddresses",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidAddressesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomersCustomerFidAddressesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCustomersCustomerFidAddressesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCustomersCustomerFidEmails adds an email address to a customer
*/
func (a *Client) PostCustomersCustomerFidEmails(params *PostCustomersCustomerFidEmailsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidEmailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidEmailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidEmails",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/emails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidEmailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomersCustomerFidEmailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCustomersCustomerFidEmailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCustomersCustomerFidInvoicesInvoiceFidCreditNote adds a credit note to a customers invoice
*/
func (a *Client) PostCustomersCustomerFidInvoicesInvoiceFidCreditNote(params *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidInvoicesInvoiceFidCreditNote",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/invoices/{invoiceFid}/creditNote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCustomersCustomerFidPhones adds a phone number to a customer
*/
func (a *Client) PostCustomersCustomerFidPhones(params *PostCustomersCustomerFidPhonesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidPhonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidPhonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidPhones",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/phones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidPhonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomersCustomerFidPhonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCustomersCustomerFidPhonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostEntitiesEntityFidEvents triggers a new event
*/
func (a *Client) PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEntitiesEntityFidEvents",
		Method:             "POST",
		PathPattern:        "/entities/{entityFid}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEntitiesEntityFidEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEntitiesEntityFidEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEntitiesEntityFidEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBrands yours brand

  Retrieve a list of all the brands within your Fortifi account

*/
func (a *Client) GetBrands(params *GetBrandsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBrandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrandsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBrands",
		Method:             "GET",
		PathPattern:        "/brands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBrandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBrandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBrandsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetMe currents user

  Retrieve information about the current connected user (you)

*/
func (a *Client) GetMe(params *GetMeParams, authInfo runtime.ClientAuthInfoWriter) (*GetMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMe",
		Method:             "GET",
		PathPattern:        "/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOrganisation currents organisation

  Retrieve information about the current organisation

*/
func (a *Client) GetOrganisation(params *GetOrganisationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganisationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganisationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganisation",
		Method:             "GET",
		PathPattern:        "/organisation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganisationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganisationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrganisationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetServiceAuthToken verifies service account get access token

  User service account credentials to retrieve an API token

*/
func (a *Client) GetServiceAuthToken(params *GetServiceAuthTokenParams) (*GetServiceAuthTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceAuthTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceAuthToken",
		Method:             "POST",
		PathPattern:        "/svcauth/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceAuthTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceAuthTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServiceAuthTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetVersion currents version

  Retrieve the current version of the Fortifi api

*/
func (a *Client) GetVersion(params *GetVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
