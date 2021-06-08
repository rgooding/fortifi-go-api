// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new marketing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for marketing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePublishersPublisherFid(params *DeletePublishersPublisherFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePublishersPublisherFidOK, error)

	DeletePublishersPublisherFidApproved(params *DeletePublishersPublisherFidApprovedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePublishersPublisherFidApprovedOK, error)

	GetPublishers(params *GetPublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersOK, error)

	GetPublishersPublisherFid(params *GetPublishersPublisherFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidOK, error)

	GetPublishersPublisherFidCampaigns(params *GetPublishersPublisherFidCampaignsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidCampaignsOK, error)

	GetPublishersPublisherFidCampaignsPublisherCampaignFid(params *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidCampaignsPublisherCampaignFidOK, error)

	GetVisitorsVisitorID(params *GetVisitorsVisitorIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVisitorsVisitorIDOK, error)

	GetVisitorsVisitorIDPixels(params *GetVisitorsVisitorIDPixelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVisitorsVisitorIDPixelsOK, error)

	PostPublishers(params *PostPublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPublishersOK, error)

	PostPublishersPublisherFidCampaigns(params *PostPublishersPublisherFidCampaignsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPublishersPublisherFidCampaignsOK, error)

	PostVisitorsVisitorIDActionsActionKey(params *PostVisitorsVisitorIDActionsActionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisitorsVisitorIDActionsActionKeyOK, error)

	PostVisitorsVisitorIDActionsActionKeyReverse(params *PostVisitorsVisitorIDActionsActionKeyReverseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisitorsVisitorIDActionsActionKeyReverseOK, error)

	PutPublishersPublisherFidApproved(params *PutPublishersPublisherFidApprovedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidApprovedOK, error)

	PutPublishersPublisherFidDisable(params *PutPublishersPublisherFidDisableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidDisableOK, error)

	PutPublishersPublisherFidEnable(params *PutPublishersPublisherFidEnableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidEnableOK, error)

	PutPublishersPublisherFidRestore(params *PutPublishersPublisherFidRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidRestoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePublishersPublisherFid deletes a publisher
*/
func (a *Client) DeletePublishersPublisherFid(params *DeletePublishersPublisherFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePublishersPublisherFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublishersPublisherFidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePublishersPublisherFid",
		Method:             "DELETE",
		PathPattern:        "/publishers/{publisherFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePublishersPublisherFidReader{formats: a.formats},
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
	success, ok := result.(*DeletePublishersPublisherFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePublishersPublisherFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePublishersPublisherFidApproved removes approved status on an publisher
*/
func (a *Client) DeletePublishersPublisherFidApproved(params *DeletePublishersPublisherFidApprovedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePublishersPublisherFidApprovedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublishersPublisherFidApprovedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePublishersPublisherFidApproved",
		Method:             "DELETE",
		PathPattern:        "/publishers/{publisherFid}/approved",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePublishersPublisherFidApprovedReader{formats: a.formats},
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
	success, ok := result.(*DeletePublishersPublisherFidApprovedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePublishersPublisherFidApprovedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPublishers lists publishers
*/
func (a *Client) GetPublishers(params *GetPublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublishersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPublishers",
		Method:             "GET",
		PathPattern:        "/publishers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublishersReader{formats: a.formats},
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
	success, ok := result.(*GetPublishersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPublishersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPublishersPublisherFid retrieves a publisher
*/
func (a *Client) GetPublishersPublisherFid(params *GetPublishersPublisherFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublishersPublisherFidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPublishersPublisherFid",
		Method:             "GET",
		PathPattern:        "/publishers/{publisherFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublishersPublisherFidReader{formats: a.formats},
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
	success, ok := result.(*GetPublishersPublisherFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPublishersPublisherFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPublishersPublisherFidCampaigns lists campaigns
*/
func (a *Client) GetPublishersPublisherFidCampaigns(params *GetPublishersPublisherFidCampaignsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidCampaignsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublishersPublisherFidCampaignsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPublishersPublisherFidCampaigns",
		Method:             "GET",
		PathPattern:        "/publishers/{publisherFid}/campaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublishersPublisherFidCampaignsReader{formats: a.formats},
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
	success, ok := result.(*GetPublishersPublisherFidCampaignsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPublishersPublisherFidCampaignsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPublishersPublisherFidCampaignsPublisherCampaignFid retrieves a publisher campaign
*/
func (a *Client) GetPublishersPublisherFidCampaignsPublisherCampaignFid(params *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPublishersPublisherFidCampaignsPublisherCampaignFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPublishersPublisherFidCampaignsPublisherCampaignFid",
		Method:             "GET",
		PathPattern:        "/publishers/{publisherFid}/campaigns/{publisherCampaignFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublishersPublisherFidCampaignsPublisherCampaignFidReader{formats: a.formats},
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
	success, ok := result.(*GetPublishersPublisherFidCampaignsPublisherCampaignFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPublishersPublisherFidCampaignsPublisherCampaignFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetVisitorsVisitorID retrieves information about a visitor

  This call will return information related to how a visitor arrived
*/
func (a *Client) GetVisitorsVisitorID(params *GetVisitorsVisitorIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVisitorsVisitorIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVisitorsVisitorIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVisitorsVisitorID",
		Method:             "GET",
		PathPattern:        "/visitors/{visitorId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVisitorsVisitorIDReader{formats: a.formats},
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
	success, ok := result.(*GetVisitorsVisitorIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVisitorsVisitorIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetVisitorsVisitorIDPixels retrieves pending pixels for this visitor

  This call will release pixels from the pending queue on read
*/
func (a *Client) GetVisitorsVisitorIDPixels(params *GetVisitorsVisitorIDPixelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVisitorsVisitorIDPixelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVisitorsVisitorIDPixelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVisitorsVisitorIDPixels",
		Method:             "GET",
		PathPattern:        "/visitors/{visitorId}/pixels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVisitorsVisitorIDPixelsReader{formats: a.formats},
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
	success, ok := result.(*GetVisitorsVisitorIDPixelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVisitorsVisitorIDPixelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostPublishers creates a new publisher
*/
func (a *Client) PostPublishers(params *PostPublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPublishersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublishersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostPublishers",
		Method:             "POST",
		PathPattern:        "/publishers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPublishersReader{formats: a.formats},
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
	success, ok := result.(*PostPublishersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostPublishersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostPublishersPublisherFidCampaigns creates a new publisher campaign
*/
func (a *Client) PostPublishersPublisherFidCampaigns(params *PostPublishersPublisherFidCampaignsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPublishersPublisherFidCampaignsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublishersPublisherFidCampaignsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostPublishersPublisherFidCampaigns",
		Method:             "POST",
		PathPattern:        "/publishers/{publisherFid}/campaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPublishersPublisherFidCampaignsReader{formats: a.formats},
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
	success, ok := result.(*PostPublishersPublisherFidCampaignsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostPublishersPublisherFidCampaignsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostVisitorsVisitorIDActionsActionKey tracks an action

  Track an action such as a lead or acquisition

*/
func (a *Client) PostVisitorsVisitorIDActionsActionKey(params *PostVisitorsVisitorIDActionsActionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisitorsVisitorIDActionsActionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVisitorsVisitorIDActionsActionKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostVisitorsVisitorIDActionsActionKey",
		Method:             "POST",
		PathPattern:        "/visitors/{visitorId}/actions/{actionKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVisitorsVisitorIDActionsActionKeyReader{formats: a.formats},
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
	success, ok := result.(*PostVisitorsVisitorIDActionsActionKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVisitorsVisitorIDActionsActionKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostVisitorsVisitorIDActionsActionKeyReverse reverses a previously tracked action

  When an action has been reversed, e.g. cancelled, refunded

*/
func (a *Client) PostVisitorsVisitorIDActionsActionKeyReverse(params *PostVisitorsVisitorIDActionsActionKeyReverseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisitorsVisitorIDActionsActionKeyReverseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVisitorsVisitorIDActionsActionKeyReverseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostVisitorsVisitorIDActionsActionKeyReverse",
		Method:             "POST",
		PathPattern:        "/visitors/{visitorId}/actions/{actionKey}/reverse",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVisitorsVisitorIDActionsActionKeyReverseReader{formats: a.formats},
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
	success, ok := result.(*PostVisitorsVisitorIDActionsActionKeyReverseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVisitorsVisitorIDActionsActionKeyReverseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutPublishersPublisherFidApproved sets approved status on an publisher
*/
func (a *Client) PutPublishersPublisherFidApproved(params *PutPublishersPublisherFidApprovedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidApprovedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPublishersPublisherFidApprovedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutPublishersPublisherFidApproved",
		Method:             "PUT",
		PathPattern:        "/publishers/{publisherFid}/approved",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPublishersPublisherFidApprovedReader{formats: a.formats},
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
	success, ok := result.(*PutPublishersPublisherFidApprovedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutPublishersPublisherFidApprovedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutPublishersPublisherFidDisable disables a publisher
*/
func (a *Client) PutPublishersPublisherFidDisable(params *PutPublishersPublisherFidDisableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidDisableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPublishersPublisherFidDisableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutPublishersPublisherFidDisable",
		Method:             "PUT",
		PathPattern:        "/publishers/{publisherFid}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPublishersPublisherFidDisableReader{formats: a.formats},
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
	success, ok := result.(*PutPublishersPublisherFidDisableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutPublishersPublisherFidDisableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutPublishersPublisherFidEnable enables a publisher
*/
func (a *Client) PutPublishersPublisherFidEnable(params *PutPublishersPublisherFidEnableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidEnableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPublishersPublisherFidEnableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutPublishersPublisherFidEnable",
		Method:             "PUT",
		PathPattern:        "/publishers/{publisherFid}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPublishersPublisherFidEnableReader{formats: a.formats},
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
	success, ok := result.(*PutPublishersPublisherFidEnableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutPublishersPublisherFidEnableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutPublishersPublisherFidRestore restores a publisher
*/
func (a *Client) PutPublishersPublisherFidRestore(params *PutPublishersPublisherFidRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutPublishersPublisherFidRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPublishersPublisherFidRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutPublishersPublisherFidRestore",
		Method:             "PUT",
		PathPattern:        "/publishers/{publisherFid}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPublishersPublisherFidRestoreReader{formats: a.formats},
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
	success, ok := result.(*PutPublishersPublisherFidRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutPublishersPublisherFidRestoreDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
