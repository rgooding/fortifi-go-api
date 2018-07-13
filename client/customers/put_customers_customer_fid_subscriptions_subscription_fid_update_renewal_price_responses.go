// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPrice structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK handles this case with default header values.

Subscription price updated
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK struct {
	Payload *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid update renewal price default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPrice default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody put customers customer fid subscriptions subscription fid update renewal price o k body
swagger:model PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) UnmarshalJSON(raw []byte) error {
	// PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0
	var putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0

	// PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO0)

	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1.Data = o.Data

	jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1, errPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1 := swag.WriteJSON(dataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1)
	if errPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1 != nil {
		return nil, errPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put customers customer fid subscriptions subscription fid update renewal price o k body
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
