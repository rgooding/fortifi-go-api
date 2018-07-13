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

// PostCustomersCustomerFidPaymentsPaymentFidChargebackReader is a Reader for the PostCustomersCustomerFidPaymentsPaymentFidChargeback structure.
type PostCustomersCustomerFidPaymentsPaymentFidChargebackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCustomersCustomerFidPaymentsPaymentFidChargebackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCustomersCustomerFidPaymentsPaymentFidChargebackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackOK creates a PostCustomersCustomerFidPaymentsPaymentFidChargebackOK with default headers values
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackOK() *PostCustomersCustomerFidPaymentsPaymentFidChargebackOK {
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackOK{}
}

/*PostCustomersCustomerFidPaymentsPaymentFidChargebackOK handles this case with default header values.

Chargeback Opened
*/
type PostCustomersCustomerFidPaymentsPaymentFidChargebackOK struct {
	Payload *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody
}

func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/payments/{paymentFid}/chargeback][%d] postCustomersCustomerFidPaymentsPaymentFidChargebackOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackDefault creates a PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault with default headers values
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackDefault(code int) *PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault {
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid payments payment fid chargeback default response
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/payments/{paymentFid}/chargeback][%d] PostCustomersCustomerFidPaymentsPaymentFidChargeback default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody post customers customer fid payments payment fid chargeback o k body
swagger:model PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody
*/
type PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0
	var postCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0

	// PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1
	var dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO0)

	var dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1, errPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1)
	if errPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid payments payment fid chargeback o k body
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidPaymentsPaymentFidChargebackOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
