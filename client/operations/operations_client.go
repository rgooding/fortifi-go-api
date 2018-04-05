// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCustomersCustomerFidPaymentMethodsCardsCardFid delete customers customer fid payment methods cards card fid API
*/
func (a *Client) DeleteCustomersCustomerFidPaymentMethodsCardsCardFid(params *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCustomersCustomerFidPaymentMethodsCardsCardFid",
		Method:             "DELETE",
		PathPattern:        "/customers/{customerFid}/paymentMethods/cards/{cardFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK), nil

}

/*
GetCustomersCustomerFidPaymentAccounts lists customers payment accounts
*/
func (a *Client) GetCustomersCustomerFidPaymentAccounts(params *GetCustomersCustomerFidPaymentAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomersCustomerFidPaymentAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomersCustomerFidPaymentAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomersCustomerFidPaymentAccounts",
		Method:             "GET",
		PathPattern:        "/customers/{customerFid}/paymentAccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomersCustomerFidPaymentAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomersCustomerFidPaymentAccountsOK), nil

}

/*
GetCustomersCustomerFidPaymentMethodsCards lists customers card payment methods
*/
func (a *Client) GetCustomersCustomerFidPaymentMethodsCards(params *GetCustomersCustomerFidPaymentMethodsCardsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomersCustomerFidPaymentMethodsCardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomersCustomerFidPaymentMethodsCardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomersCustomerFidPaymentMethodsCards",
		Method:             "GET",
		PathPattern:        "/customers/{customerFid}/paymentMethods/cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomersCustomerFidPaymentMethodsCardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomersCustomerFidPaymentMethodsCardsOK), nil

}

/*
GetCustomersCustomerFidTickets gets support tickets for customer
*/
func (a *Client) GetCustomersCustomerFidTickets(params *GetCustomersCustomerFidTicketsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomersCustomerFidTicketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomersCustomerFidTicketsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomersCustomerFidTickets",
		Method:             "GET",
		PathPattern:        "/customers/{customerFid}/tickets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomersCustomerFidTicketsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomersCustomerFidTicketsOK), nil

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
	return result.(*GetPayCoinbaseOK), nil

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
	return result.(*GetPayPublicKeyOK), nil

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
	return result.(*PostCustomersCustomerFidEmailsOK), nil

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
	return result.(*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteOK), nil

}

/*
PostCustomersCustomerFidPaymentMethodsCards adds a new card
*/
func (a *Client) PostCustomersCustomerFidPaymentMethodsCards(params *PostCustomersCustomerFidPaymentMethodsCardsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidPaymentMethodsCardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidPaymentMethodsCardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidPaymentMethodsCards",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/paymentMethods/cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidPaymentMethodsCardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomersCustomerFidPaymentMethodsCardsOK), nil

}

/*
PostCustomersCustomerFidPaymentMethodsPaypalComplete completes a paypal agreement created with initialise
*/
func (a *Client) PostCustomersCustomerFidPaymentMethodsPaypalComplete(params *PostCustomersCustomerFidPaymentMethodsPaypalCompleteParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidPaymentMethodsPaypalCompleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidPaymentMethodsPaypalCompleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidPaymentMethodsPaypalComplete",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/paymentMethods/paypal/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidPaymentMethodsPaypalCompleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomersCustomerFidPaymentMethodsPaypalCompleteOK), nil

}

/*
PostCustomersCustomerFidPaymentMethodsPaypalInitialise initialises a new paypal agreement for existing subscriptions
*/
func (a *Client) PostCustomersCustomerFidPaymentMethodsPaypalInitialise(params *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidPaymentMethodsPaypalInitialise",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/paymentMethods/paypal/initialise",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK), nil

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
	return result.(*PostCustomersCustomerFidPhonesOK), nil

}

/*
PostCustomersCustomerFidTickets creates a new support ticket
*/
func (a *Client) PostCustomersCustomerFidTickets(params *PostCustomersCustomerFidTicketsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomersCustomerFidTicketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomersCustomerFidTicketsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomersCustomerFidTickets",
		Method:             "POST",
		PathPattern:        "/customers/{customerFid}/tickets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomersCustomerFidTicketsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomersCustomerFidTicketsOK), nil

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
	return result.(*PostEntitiesEntityFidEventsOK), nil

}

/*
PutCustomersCustomerFidPaymentMethodsCardsCardFid updates a card
*/
func (a *Client) PutCustomersCustomerFidPaymentMethodsCardsCardFid(params *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomersCustomerFidPaymentMethodsCardsCardFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCustomersCustomerFidPaymentMethodsCardsCardFid",
		Method:             "PUT",
		PathPattern:        "/customers/{customerFid}/paymentMethods/cards/{cardFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutCustomersCustomerFidPaymentMethodsCardsCardFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCustomersCustomerFidPaymentMethodsCardsCardFidOK), nil

}

/*
PutMessengerDeliveriesDeliveryFidUnsubscribe unsubscribes an email based on the delivery fid
*/
func (a *Client) PutMessengerDeliveriesDeliveryFidUnsubscribe(params *PutMessengerDeliveriesDeliveryFidUnsubscribeParams, authInfo runtime.ClientAuthInfoWriter) (*PutMessengerDeliveriesDeliveryFidUnsubscribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMessengerDeliveriesDeliveryFidUnsubscribeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutMessengerDeliveriesDeliveryFidUnsubscribeOK), nil

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
	return result.(*GetBrandsOK), nil

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
	return result.(*GetMeOK), nil

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
	return result.(*GetOrganisationOK), nil

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
	return result.(*GetServiceAuthTokenOK), nil

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
	return result.(*GetVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
