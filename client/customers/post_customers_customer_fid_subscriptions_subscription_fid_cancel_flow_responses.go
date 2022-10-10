// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReader is a Reader for the PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlow structure.
type PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK creates a PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK with default headers values
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK() *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK{}
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK describes a response with status code 200, with default header values.

Cancel Flow State
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK struct {
	Payload *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody
}

// IsSuccess returns true when this post customers customer fid subscriptions subscription fid cancel flow o k response has a 2xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post customers customer fid subscriptions subscription fid cancel flow o k response has a 3xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post customers customer fid subscriptions subscription fid cancel flow o k response has a 4xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post customers customer fid subscriptions subscription fid cancel flow o k response has a 5xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post customers customer fid subscriptions subscription fid cancel flow o k response a status code equal to that given
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow][%d] postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow][%d] postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) GetPayload() *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault creates a PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault with default headers values
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault(code int) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault{
		_statusCode: code,
	}
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid subscriptions subscription fid cancel flow default response
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post customers customer fid subscriptions subscription fid cancel flow default response has a 2xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post customers customer fid subscriptions subscription fid cancel flow default response has a 3xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post customers customer fid subscriptions subscription fid cancel flow default response has a 4xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post customers customer fid subscriptions subscription fid cancel flow default response has a 5xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post customers customer fid subscriptions subscription fid cancel flow default response a status code equal to that given
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow][%d] PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlow default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow][%d] PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlow default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody post customers customer fid subscriptions subscription fid cancel flow o k body
swagger:model PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody struct {
	models.Envelope

	// data
	Data *models.CancelFlowState `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0
	var postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0

	// PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1
	var dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1 struct {
		Data *models.CancelFlowState `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO0)
	var dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1 struct {
		Data *models.CancelFlowState `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1, errPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1)
	if errPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid subscriptions subscription fid cancel flow o k body
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post customers customer fid subscriptions subscription fid cancel flow o k body based on the context it is used
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
