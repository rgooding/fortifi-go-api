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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandon structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK{}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK describes a response with status code 200, with default header values.

Retention Flow State
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK struct {
	Payload *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid retention flow flow search abandon o k response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/abandon][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/abandon][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) GetPayload() *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon default response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon default response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon default response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon default response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid retention flow flow search abandon default response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid retention flow flow search abandon default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/abandon][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandon default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/abandon][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandon default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody put customers customer fid subscriptions subscription fid retention flow flow search abandon o k body
swagger:model PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody struct {
	models.Envelope

	// data
	Data *models.RetentionFlowState `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) UnmarshalJSON(raw []byte) error {
	// PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0
	var putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0

	// PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1 struct {
		Data *models.RetentionFlowState `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO0)
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1 struct {
		Data *models.RetentionFlowState `json:"data,omitempty"`
	}

	dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1.Data = o.Data

	jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1, errPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1 := swag.WriteJSON(dataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1)
	if errPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1 != nil {
		return nil, errPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k body
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put customers customer fid subscriptions subscription fid retention flow flow search abandon o k body based on the context it is used
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchAbandonOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}