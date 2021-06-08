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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenew structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK{}
}

/* PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK describes a response with status code 200, with default header values.

Created new open order
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK struct {
	Payload *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/preRenew][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK  %+v", 200, o.Payload)
}
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK) GetPayload() *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault{
		_statusCode: code,
	}
}

/* PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid pre renew default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/preRenew][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenew default  %+v", o._statusCode, o.Payload)
}
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody put customers customer fid subscriptions subscription fid pre renew o k body
swagger:model PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) UnmarshalJSON(raw []byte) error {
	// PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0
	var putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0

	// PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO0)
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1.Data = o.Data

	jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1, errPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1 := swag.WriteJSON(dataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1)
	if errPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1 != nil {
		return nil, errPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put customers customer fid subscriptions subscription fid pre renew o k body
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put customers customer fid subscriptions subscription fid pre renew o k body based on the context it is used
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomersCustomerFidSubscriptionsSubscriptionFidPreRenewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
